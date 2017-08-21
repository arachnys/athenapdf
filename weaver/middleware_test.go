package main

import (
	"errors"
	"github.com/arachnys/athenapdf/weaver/converter"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestConfigMiddleware(t *testing.T) {
	r := gin.Default()
	mockConf := Config{MaxWorkers: 17}
	var ctxConf Config
	r.Use(ConfigMiddleware(mockConf))
	r.GET("/", func(c *gin.Context) {
		ctxConf = c.MustGet("config").(Config)
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res, req)
	if got, want := res.Code, http.StatusOK; got != want {
		t.Fatalf("expected response code to be %d, got %d", want, got)
	}
	if !reflect.DeepEqual(mockConf, ctxConf) {
		t.Errorf("expected config in context to be %+v, got %+v", mockConf, ctxConf)
	}
}

func TestWorkQueueMiddleware(t *testing.T) {
	r := gin.Default()
	mockWq := make(chan converter.Work, 17)
	var ctxWq chan<- converter.Work
	r.Use(WorkQueueMiddleware(mockWq))
	r.GET("/", func(c *gin.Context) {
		ctxWq = c.MustGet("queue").(chan<- converter.Work)
		go func(wq chan<- converter.Work) {
			ctxWq <- converter.Work{}
		}(ctxWq)
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res, req)
	if got, want := res.Code, http.StatusOK; got != want {
		t.Fatalf("expected response code to be %d, got %d", want, got)
	}
	select {
	case <-mockWq:
	case <-time.After(time.Second):
		t.Errorf("expected to receive work from queue before timeout")
	}
}

func TestSentryMiddleware(t *testing.T) {
	r := gin.Default()
	mockRaven := new(raven.Client)
	var ctxRaven *raven.Client
	r.Use(SentryMiddleware(mockRaven))
	r.GET("/", func(c *gin.Context) {
		ctxRaven = c.MustGet("sentry").(*raven.Client)
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res, req)
	if got, want := res.Code, http.StatusOK; got != want {
		t.Fatalf("expected response code to be %d, got %d", want, got)
	}
	if !reflect.DeepEqual(mockRaven, ctxRaven) {
		t.Errorf("expected sentry in context to be %+v, got %+v", mockRaven, ctxRaven)
	}
}

func TestStatsdMiddleware(t *testing.T) {
	r := gin.Default()
	mockStatsd := new(statsd.Client)
	var ctxStatsd *statsd.Client
	r.Use(StatsdMiddleware(mockStatsd))
	r.GET("/", func(c *gin.Context) {
		ctxStatsd = c.MustGet("statsd").(*statsd.Client)
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res, req)
	if got, want := res.Code, http.StatusOK; got != want {
		t.Fatalf("expected response code to be %d, got %d", want, got)
	}
	if !reflect.DeepEqual(mockStatsd, ctxStatsd) {
		t.Errorf("expected statsd in context to be %+v, got %+v", mockStatsd, ctxStatsd)
	}
}

func TestErrorMiddleware(t *testing.T) {
	r := gin.Default()
	r.Use(ErrorMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.Error(errors.New("test error"))
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res, req)
	if got, want := res.Code, http.StatusInternalServerError; got != want {
		t.Fatalf("expected response code to be %d, got %d", want, got)
	}
	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %+v", err)
	}
	want := "{\"error\":\"PDF conversion failed due to an internal server error\"}"
	if !reflect.DeepEqual(strings.TrimSpace(string(got)), want) {
		t.Errorf("expected response body to be %s, got %s", want, got)
	}
}

func TestAuthorizationMiddleware(t *testing.T) {
	r := gin.Default()
	r.Use(AuthorizationMiddleware("123456"))
	r.GET("/", func(c *gin.Context) {})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?auth=123456", nil)
	r.ServeHTTP(res, req)
	if got, want := res.Code, http.StatusOK; got != want {
		t.Fatalf("expected response code to be %d, got %d", want, got)
	}
}

func TestAuthorizationMiddleware_authFailure(t *testing.T) {
	r := gin.Default()
	r.Use(AuthorizationMiddleware("123456"))
	r.GET("/", func(c *gin.Context) {})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(res, req)
	if got, want := res.Code, http.StatusUnauthorized; got != want {
		t.Fatalf("expected response code to be %d, got %d", want, got)
	}
}

package http

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func loadFile(fp string) ([]byte, error) {
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	if f != nil {
		defer f.Close()
	}
	return ioutil.ReadAll(f)
}

func TestFetch(t *testing.T) {
	testComplexHtmlData, err := loadFile("testdata/complex_html.html")
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	testSimpleXmlData, err := loadFile("testdata/simple_xml.xml")
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	testComplexXmlData, err := loadFile("testdata/complex_xml.xml")
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	testCases := []struct {
		name            string
		data            []byte
		contentType     string
		wantContentType string
	}{
		{
			"simple text",
			[]byte("simple test data"),
			"text/plain; charset=\"UTF-8\"",
			"text/plain; charset=\"UTF-8\"",
		},
		{
			"simple text with no content-type header",
			[]byte("simple test data"),
			"",
			"text/plain; charset=\"UTF-8\"",
		},
		{
			"complex html with no content-type header",
			testComplexHtmlData,
			"",
			"text/html; charset=utf-8",
		},
		{
			"simple xml with no content-type header",
			testSimpleXmlData,
			"",
			"text/plain; charset=\"UTF-8\"",
		},
		{
			"complex xml with no content-type header",
			testComplexXmlData,
			"",
			"text/xml; charset=utf-8",
		},
	}

	ctx := context.Background()
	fetcher := &HTTPFetcher{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.Write(tc.data)
			}))
			defer ts.Close()

			gotReader, gotContentType, err := fetcher.Fetch(ctx, ts.URL, nil)
			if err != nil {
				t.Fatalf("failed to fetch, unexpected error: %+v", err)
			}

			if gotContentType != tc.wantContentType {
				t.Errorf("got %s; want %s", gotContentType, tc.wantContentType)
			}

			gotData, err := ioutil.ReadAll(gotReader)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}

			if !reflect.DeepEqual(gotData, tc.data) {
				t.Errorf("got %s; want %s", gotData, tc.data)
			}
		})
	}

	t.Run("basic auth", func(t *testing.T) {
		wantData := []byte("test protected page")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, p, ok := r.BasicAuth()
			if !ok || (u != "test" && p != "test") {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("unauthorized"))
				return
			}

			w.Write(wantData)
		}))
		defer ts.Close()

		u, err := url.Parse(ts.URL)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}
		u.User = url.UserPassword("test", "test")

		gotReader, _, err := fetcher.Fetch(ctx, u.String(), nil)
		if err != nil {
			t.Fatalf("failed to fetch, unexpected error: %+v", err)
		}

		gotData, err := ioutil.ReadAll(gotReader)
		if err != nil {
			t.Fatalf("unexpected error: %+v", err)
		}

		if !reflect.DeepEqual(gotData, wantData) {
			t.Errorf("got %s; want %s", gotData, wantData)
		}
	})

	t.Run("context cancellation", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Fatalf("unexpected http request, context should have been cancelled")
		}))
		defer ts.Close()

		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, _, err := fetcher.Fetch(ctx, ts.URL, nil)
		if err != context.Canceled {
			t.Fatalf("got %+v; want %+v", err, context.Canceled)
		}
	})
}

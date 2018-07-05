package main

import (
	"context"
	"github.com/DeanThompson/ginpprof"
	"github.com/arachnys/athenapdf/weaver/converter"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/contrib/sentry"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

// InitMiddleware sets up the necessary middlewares for the microservice.
// These include middlewares to establish a sane context containing access to
// the configuration, worker queue, statsd client, and Sentry client (Raven).
// The latter two are disabled in debugging mode to avoid contaminating
// production stats.
// It will also set up a middleware for catching, and handling errors thrown
// from a route.
func InitMiddleware(router *gin.Engine, conf Config) {
	// Config
	router.Use(ConfigMiddleware(conf))

	// Worker queue
	wq := converter.InitWorkers(conf.MaxWorkers, conf.MaxConversionQueue, conf.WorkerTimeout)
	router.Use(WorkQueueMiddleware(wq))

	// Statsd
	muteStatsd := gin.IsDebugging()
	if conf.Statsd.Address == "" {
		muteStatsd = true
	}
	s, err := statsd.New(
		statsd.Address(conf.Statsd.Address),
		statsd.Prefix(conf.Statsd.Prefix),
		statsd.FlushPeriod(time.Millisecond*500),
		statsd.Mute(muteStatsd),
	)
	if err != nil {
		panic(err)
	}
	router.Use(StatsdMiddleware(s))

	// Sentry (crash reporting)
	if !gin.IsDebugging() && conf.SentryDSN != "" {
		r, err := raven.New(conf.SentryDSN)
		if err != nil {
			panic(err)
		}
		router.Use(SentryMiddleware(r))
		router.Use(sentry.Recovery(r, true))
	}

	// Error handler
	router.Use(ErrorMiddleware())
}

// InitSecureRoutes creates the necessary conversion routes with a middleware
// to restrict access via an auth key (defined in the environment config).
func InitSecureRoutes(router *gin.Engine, conf Config) {
	authorized := router.Group("/")
	authorized.Use(AuthorizationMiddleware(conf.AuthKey))
	authorized.GET("/convert", convertByURLHandler)
	authorized.POST("/convert", convertByFileHandler)
}

// InitSimpleRoutes creates non-essential routes for monitoring and/or
// debugging.
func InitSimpleRoutes(router *gin.Engine, conf Config) {
	router.GET("/", indexHandler)
	router.GET("/stats", statsHandler)

	if gin.IsDebugging() {
		ginpprof.Wrapper(router)
	}
}

func StartX() {
	cmd := exec.Command("Xvfb", ":99", "-ac", "-screen", "0", "1024x768x24")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	err := cmd.Run()
	log.Fatal("Xvfb exited:", err)
}

func main() {
	router := gin.Default()
	// Get config vars from the environment
	conf := NewEnvConfig()
	InitMiddleware(router, conf)
	InitSecureRoutes(router, conf)
	InitSimpleRoutes(router, conf)

	if conf.HTTPSAddr != "" {
		if conf.TLSCertFile == "" {
			log.Fatal("No TLS cert file provided (WEAVER_TLS_CERT_FILE)")
		}

		if conf.TLSKeyFile == "" {
			log.Fatal("No TLS key file provided (WEAVER_TLS_KEY_FILE)")
		}

		go func() {
			log.Fatal(router.RunTLS(conf.HTTPSAddr, conf.TLSCertFile, conf.TLSKeyFile))
		}()
	}

	server := &http.Server{
		Addr:    conf.HTTPAddr,
		Handler: router,
	}

	go func() {
		log.Println(server.ListenAndServe())
	}()

	go StartX()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	log.Println("Received sigterm, gracefully shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error:", err)
	}

}

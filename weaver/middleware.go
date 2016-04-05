package main

import (
	"errors"
	"github.com/arachnys/athenapdf/weaver/converter"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"net/http"
)

var (
	// ErrAuthorization should be returned when the authorization key is invalid.
	ErrAuthorization = errors.New("invalid authorization key provided")
	// ErrInternalServer should be returned when a private error is returned
	// from a handler.
	ErrInternalServer = errors.New("PDF conversion failed due to an internal server error")
)

// ConfigMiddleware sets the config in the context.
func ConfigMiddleware(conf Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", conf)
	}
}

// WorkQueueMiddleware sets the work queue (write only) in the context.
func WorkQueueMiddleware(q chan<- converter.Work) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("queue", q)
	}
}

// SentryMiddleware sets the Sentry client (Raven) in the context.
func SentryMiddleware(r *raven.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("sentry", r)
	}
}

// StatsdMiddleware sets the Statsd client in the context.
func StatsdMiddleware(s *statsd.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("statsd", s)
	}
}

// ErrorMiddleware runs after all handlers have been executed, and it handles
// any errors returned from the handlers. It will return an internal server
// error with a predefined message if the last error type is not public.
// Otherwise, it will display the last error message it received, and the
// associated HTTP status code.
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastError := c.Errors.Last()
		statusCode := c.Writer.Status()

		if lastError != nil {
			// Log all errors
			log.Println("captured errors:")
			log.Printf("%+v\n", c.Errors)

			// Public errors
			if lastError.IsType(gin.ErrorTypePublic) {
				c.JSON(statusCode, gin.H{
					"error": lastError.Error(),
				})
				return
			}

			// Private errors
			c.JSON(500, gin.H{
				"error": ErrInternalServer.Error(),
			})
		}
	}
}

// AuthorizationMiddleware is a simple authorization middleware which matches
// an authentication key, provided via a query parameter, against a defined
// authentication key in the environment config.
func AuthorizationMiddleware(k string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("auth") != k {
			c.AbortWithError(http.StatusUnauthorized, ErrAuthorization).SetType(gin.ErrorTypePublic)
		}

		c.Next()
	}
}

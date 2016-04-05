package main

import (
	"errors"
	"github.com/arachnys/athenapdf/weaver/converter"
	"github.com/arachnys/athenapdf/weaver/converter/athenapdf"
	"github.com/arachnys/athenapdf/weaver/converter/cloudconvert"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
	"log"
	"net/http"
	"runtime"
)

var (
	// ErrURLInvalid should be returned when a conversion URL is invalid.
	ErrURLInvalid = errors.New("invalid URL provided")
)

// indexHandler returns a JSON string indicating that the microservice is online.
// It does not actually check if conversions are working. It is nevertheless,
// used for monitoring.
func indexHandler(c *gin.Context) {
	// We can do better than this...
	c.JSON(http.StatusOK, gin.H{"status": "online"})
}

// statsHandler returns a JSON string containing the number of running
// Goroutines, and pending jobs in the work queue.
func statsHandler(c *gin.Context) {
	q := c.MustGet("queue").(chan<- converter.Work)
	c.JSON(http.StatusOK, gin.H{
		"goroutines": runtime.NumGoroutine(),
		"pending":    len(q),
	})
}

// convertByURLHandler is the main v1 API handler for converting a HTML to a PDF
// via a GET request. It can either return a JSON string indicating that the
// output of the conversion has been uploaded or it can return the output of
// the conversion to the client (raw bytes).
func convertByURLHandler(c *gin.Context) {
	conf := c.MustGet("config").(Config)
	wq := c.MustGet("queue").(chan<- converter.Work)
	s := c.MustGet("statsd").(*statsd.Client)
	r, ravenOk := c.Get("sentry")

	t := s.NewTiming()

	url := c.Query("url")
	_, aggressive := c.GetQuery("aggressive")
	awsConf := converter.AWSS3{
		c.Query("aws_region"),
		c.Query("aws_id"),
		c.Query("aws_secret"),
		c.Query("s3_bucket"),
		c.Query("s3_key"),
	}

	if url == "" {
		c.AbortWithError(http.StatusBadRequest, ErrURLInvalid).SetType(gin.ErrorTypePublic)
		s.Increment("invalid_url")
		return
	}

	var conversion converter.Converter
	var work converter.Work
	attempts := 0

	baseConversion := converter.Conversion{url}
	uploadConversion := converter.UploadConversion{baseConversion, awsConf}

StartConversion:
	conversion = athenapdf.AthenaPDF{uploadConversion, conf.AthenaCMD, aggressive}
	if attempts != 0 {
		cc := cloudconvert.Client{conf.CloudConvert.APIUrl, conf.CloudConvert.APIKey}
		conversion = cloudconvert.CloudConvert{uploadConversion, cc}
	}
	work = converter.NewWork(wq, conversion)

	select {
	case <-c.Writer.CloseNotify():
		work.Cancel()
	case <-work.Uploaded():
		t.Send("conversion_duration")
		s.Increment("success")
		c.JSON(200, gin.H{"status": "uploaded"})
	case out := <-work.Success():
		t.Send("conversion_duration")
		s.Increment("success")
		c.Data(200, "application/pdf", out)
	case err := <-work.Error():
		// log.Println(err)

		// Log, and stats collection
		if err == converter.ErrConversionTimeout {
			s.Increment("conversion_timeout")
		} else if _, awsError := err.(awserr.Error); awsError {
			s.Increment("s3_upload_error")
			if ravenOk {
				r.(*raven.Client).CaptureError(err, map[string]string{"url": url})
			}
		} else {
			s.Increment("conversion_error")
			if ravenOk {
				r.(*raven.Client).CaptureError(err, map[string]string{"url": url})
			}
		}

		if attempts == 0 && conf.ConversionFallback {
			s.Increment("cloudconvert")
			log.Println("falling back to CloudConvert...")
			attempts++
			goto StartConversion
		}

		s.Increment("conversion_failed")

		if err == converter.ErrConversionTimeout {
			c.AbortWithError(http.StatusGatewayTimeout, converter.ErrConversionTimeout).SetType(gin.ErrorTypePublic)
			return
		}

		c.Error(err)
	}
}

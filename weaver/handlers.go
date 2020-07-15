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
	"os"
	"runtime"
	"time"
)

var (
	// ErrURLInvalid should be returned when a conversion URL is invalid.
	ErrURLInvalid = errors.New("invalid URL provided")
	// ErrFileInvalid should be returned when a conversion file is invalid.
	ErrFileInvalid = errors.New("invalid file provided")
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

func conversionHandler(c *gin.Context, source converter.ConversionSource) {
	// GC if converting temporary file
	if source.IsLocal {
		defer os.Remove(source.URI)
	}

	_, aggressive := c.GetQuery("aggressive")
	_, waitForStatus := c.GetQuery("waitForStatus")

	conf := c.MustGet("config").(Config)
	wq := c.MustGet("queue").(chan<- converter.Work)
	s := c.MustGet("statsd").(*statsd.Client)
	r, ravenOk := c.Get("sentry")

	t := s.NewTiming()

	awsConf := converter.AWSS3{
		c.Query("aws_region"),
		c.Query("aws_id"),
		c.Query("aws_secret"),
		c.Query("s3_bucket"),
		c.Query("s3_key"),
		c.Query("s3_acl"),
	}

	var conversion converter.Converter
	var work converter.Work
	attempts := 0

	baseConversion := converter.Conversion{}
	uploadConversion := converter.UploadConversion{baseConversion, awsConf}

StartConversion:
	conversion = athenapdf.AthenaPDF{uploadConversion, conf.AthenaCMD, aggressive, waitForStatus}
	if attempts != 0 {
		cc := cloudconvert.Client{
			conf.CloudConvert.APIUrl,
			conf.CloudConvert.APIKey,
			time.Second * time.Duration(conf.WorkerTimeout+5),
		}
		conversion = cloudconvert.CloudConvert{uploadConversion, cc}
	}
	work = converter.NewWork(wq, conversion, source)

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
				r.(*raven.Client).CaptureError(err, map[string]string{"url": source.GetActualURI()})
			}
		} else {
			s.Increment("conversion_error")
			if ravenOk {
				r.(*raven.Client).CaptureError(err, map[string]string{"url": source.GetActualURI()})
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

// convertByURLHandler is the main v1 API handler for converting a HTML to a PDF
// via a GET request. It can either return a JSON string indicating that the
// output of the conversion has been uploaded or it can return the output of
// the conversion to the client (raw bytes).
func convertByURLHandler(c *gin.Context) {
	s := c.MustGet("statsd").(*statsd.Client)
	r, ravenOk := c.Get("sentry")

	url := c.Query("url")
	if url == "" {
		c.AbortWithError(http.StatusBadRequest, ErrURLInvalid).SetType(gin.ErrorTypePublic)
		s.Increment("invalid_url")
		return
	}

	ext := c.Query("ext")

	source, err := converter.NewConversionSource(url, nil, ext)
	if err != nil {
		s.Increment("conversion_error")
		if ravenOk {
			r.(*raven.Client).CaptureError(err, map[string]string{"url": url})
		}
		c.Error(err)
		return
	}

	conversionHandler(c, *source)
}

func convertByFileHandler(c *gin.Context) {
	s := c.MustGet("statsd").(*statsd.Client)
	r, ravenOk := c.Get("sentry")

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, ErrFileInvalid).SetType(gin.ErrorTypePublic)
		s.Increment("invalid_file")
		return
	}

	ext := c.Query("ext")

	source, err := converter.NewConversionSource("", file, ext)
	if err != nil {
		s.Increment("conversion_error")
		if ravenOk {
			r.(*raven.Client).CaptureError(err, map[string]string{"url": header.Filename})
		}
		c.Error(err)
		return
	}

	conversionHandler(c, *source)
}

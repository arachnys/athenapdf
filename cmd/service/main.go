package main

import (
	"fmt"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"gopkg.in/alecthomas/kingpin.v2"
	stdlog "log"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/arachnys/athenapdf/pkg/processor"

	_ "github.com/arachnys/athenapdf/pkg/converter/athenapdf"
	_ "github.com/arachnys/athenapdf/pkg/converter/cloudconvert"
	_ "github.com/arachnys/athenapdf/pkg/converter/libreoffice"
	_ "github.com/arachnys/athenapdf/pkg/fetcher/http"
	_ "github.com/arachnys/athenapdf/pkg/uploader/s3"
)

const (
	appName        = "weaver"
	appVersion     = "3.0.0"
	appDescription = "microservice for handling (M)HTML to PDF conversion processes"
)

var (
	app         = kingpin.New(appName, appDescription).Version(appVersion)
	debug       = app.Flag("debug", "enable debug mode / verbose logging").Short('D').Bool()
	concurrency = app.Flag("concurrency", "maximum number of conversions to process in a given time").Short('C').Int()
	host        = app.Flag("host", "host:port for the microservice to bind").Short('H').Default(":8080").String()
	jaegerHost  = app.Flag("jaeger", "host:port of a `jaeger-agent` for the reporter to send opentracing spans").String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	startedAt := time.Now()

	// Base logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.With(logger, "transport", "HTTP")
	}

	// Instrumentation
	cfg := jaegercfg.Configuration{
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           *debug,
			LocalAgentHostPort: *jaegerHost,
		},
	}
	if *debug {
		cfg.Sampler = &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		}
	}
	tracer, closer, err := cfg.New(
		appName,
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Tag("version", appVersion),
	)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer closer.Close()

	// Process manager
	manager := processor.NewManager(*concurrency, 0, 0)

	var svc PDFService
	svc = pdfService{tracer}
	svc = errorMiddleware{tracer, svc}
	svc = managerMiddleware{manager, svc}

	m := http.NewServeMux()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	m.Handle("/process", httptransport.NewServer(
		opentracing.TraceServer(tracer, "process")(processEndpoint(svc)),
		decodeProcessRequest,
		encodeProcessResponse,
		options...,
	))

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s/%s (%d)", appName, appVersion, startedAt.Unix())))
	})

	logger.Log("version", appVersion, "debug", *debug, "addr", *host)
	stdlog.Fatalln(http.ListenAndServe(*host, m))
}

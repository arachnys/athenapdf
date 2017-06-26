package main

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	stdlog "log"
	"net/http"
	"os"

	"github.com/arachnys/athenapdf/pkg/process"
	"github.com/arachnys/athenapdf/pkg/proto"

	_ "github.com/arachnys/athenapdf/pkg/converter/athenapdf"
	_ "github.com/arachnys/athenapdf/pkg/fetcher/http"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Response struct {
	Error   string `json:"error,omitempty"`
	Success string `json:"success,omitempty"`
}

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
		logger = log.With(logger, "transport", "HTTP")
	}

	m := http.NewServeMux()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	m.Handle("/process", httptransport.NewServer(
		processEndpoint,
		decodeProcessRequest,
		encodeProcessResponse,
		options...,
	))

	stdlog.Fatalln(http.ListenAndServe(":8080", m))
}

// TODO:
// - Built-in retries / fallbacks
// - Queue'ing / throttle
// - Instrumentation
func processEndpoint(ctx context.Context, request interface{}) (interface{}, error) {
	p := request.(proto.Process)

	return process.Process(ctx, &p)
}

func decodeProcessRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var process proto.Process

	switch r.Method {
	case "GET":
		query := r.URL.Query()
		if query.Get("uri") == "" {
			return nil, errors.Errorf("no uri specified")
		}
		process = proto.Process{
			Converter: query.Get("converter"),
			Conversion: &proto.Conversion{
				Uri:      query.Get("uri"),
				MimeType: query.Get("mime_type"),
			},
			Uploader: query.Get("uploader"),
			Fetcher:  query.Get("fetcher"),
		}
	case "POST":
		// TODO
	default:
		return nil, errors.Errorf("http method `%s` is not supported", r.Method)
	}

	return process, nil
}

func encodeProcessResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if v, ok := response.(Response); ok {
		return json.NewEncoder(w).Encode(v)
	}

	if v, ok := response.(io.Reader); ok {
		w.Header().Set("content-type", "application/pdf")
		if _, err := io.Copy(w, v); err != nil {
			return err
		}
	}

	return nil
}

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&Response{Error: err.Error()})
}

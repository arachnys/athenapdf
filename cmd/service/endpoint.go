package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"

	"github.com/arachnys/athenapdf/pkg/mime"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/uri"
)

type Response struct {
	Error   string `json:"error,omitempty"`
	Success string `json:"success,omitempty"`
}

type Request struct {
	File    io.Reader
	Process *proto.Process
}

func processEndpoint(svc PDFService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var res interface{}

		req := request.(Request)
		process := req.Process

		if req.Process.GetConversion().GetUri() == "" && req.File == nil {
			return nil, errors.New("no uri or file provided")
		} else if req.File != nil {
			tmp, err := ioutil.TempFile("", "athenapdf-request")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			if _, err := io.Copy(tmp, req.File); err != nil {
				return nil, errors.WithStack(err)
			}
			defer os.Remove(tmp.Name())

			if process.GetConversion() == nil {
				process.Conversion = &proto.Conversion{}
			}

			localURI, err := uri.ToLocal(tmp.Name())
			if err != nil {
				return nil, err
			}

			process.Conversion.Uri = localURI
			mt, err := mime.TypeFromFile(tmp.Name())
			if err != nil {
				return nil, errors.WithStack(err)
			}
			process.Conversion.MimeType = mt
		} else if uri.IsLocal(process.GetConversion().GetUri()) {
			return nil, errors.New("local conversions are not allowed")
		}

		res, uploaded, err := svc.Process(ctx, process)
		if err != nil {
			return nil, err
		}

		if uploaded {
			res = Response{Success: fmt.Sprintf("uploaded to %s", process.GetFetcher())}
		}

		return res, nil
	}
}

func decodeProcessRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request Request
	var process proto.Process

	switch r.Method {
	case "GET":
		query := r.URL.Query()
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
		req := r.FormValue("process")

		if err := jsonpb.UnmarshalString(req, &process); err != nil {
			return nil, errors.WithStack(err)
		}

		f, _, err := r.FormFile("file")
		if err != nil && err != http.ErrMissingFile {
			return nil, errors.WithStack(err)
		}

		if f != nil {
			var file bytes.Buffer
			if _, err := io.Copy(&file, f); err != nil {
				return nil, errors.WithStack(err)
			}
			defer f.Close()

			request.File = &file
			// TODO: use `header` for MIME type
		}
	default:
		return nil, errors.Errorf("http method `%s` is not supported", r.Method)
	}

	request.Process = &process

	return request, nil
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

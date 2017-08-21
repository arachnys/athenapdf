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

		var input io.Reader = req.File
		var mimeType string
		var err error

		if req.Process.GetFetcher() != "" {
			input, mimeType, err = svc.Fetch(ctx, req.Process)
			if err != nil {
				return nil, err
			}
		}

		if input != nil {
			tmp, err := ioutil.TempFile("", "athenapdf-process-endpoint")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			if _, err := io.Copy(tmp, input); err != nil {
				return nil, errors.WithStack(err)
			}
			defer os.Remove(tmp.Name())

			localURI, err := uri.ToLocal(tmp.Name())
			if err != nil {
				return nil, err
			}

			if req.Process.GetConversion() == nil {
				req.Process.Conversion = &proto.Conversion{}
			}

			req.Process.Conversion.Uri = localURI
			if req.Process.GetConversion().GetMimeType() == "" {
				if mimeType == "" {
					mimeType, err = mime.TypeFromFile(tmp.Name())
					if err != nil {
						return nil, errors.WithStack(err)
					}
				}
				req.Process.Conversion.MimeType = mimeType
			}
		} else if uri.IsLocal(req.Process.GetConversion().GetUri()) {
			return nil, errors.New("local conversions are not allowed")
		}

		r, err := svc.Convert(ctx, req.Process)
		if err != nil {
			return nil, err
		}

		if req.Process.GetUploader() != "" {
			if err := svc.Upload(ctx, req.Process, r); err != nil {
				return nil, err
			}
			return Response{
				Success: fmt.Sprintf("uploaded to %s", req.Process.GetFetcher()),
			}, nil
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

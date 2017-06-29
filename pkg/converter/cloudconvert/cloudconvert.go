package cloudconvert

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/arachnys/athenapdf/pkg/config"
	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/mime"
	"github.com/arachnys/athenapdf/pkg/proto"
)

const (
	converterName = "cloudconvert"
)

type CloudConvert struct{}

func init() {
	converter.Register(converterName, &CloudConvert{})
}

func (_ *CloudConvert) Convert(ctx context.Context, req *proto.Conversion, opts map[string]*proto.Option) (io.Reader, error) {
	conf := config.MustGet(converterName, opts)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	params := map[string]string{
		"apikey":       conf("api_key"),
		"download":     "inline",
		"filename":     "arachnys-athenapdf-pkg-converter-cloudconvert.html",
		"inputformat":  mime.ToExtension(req.GetMimeType()),
		"outputformat": "pdf",
	}

	if converter.IsLocal(req) {
		f, err := os.Open(req.GetUri())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if f != nil {
			defer f.Close()
		}

		fileName := filepath.Base(req.GetUri())
		part, err := w.CreateFormFile("file", fileName)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		_, err = io.Copy(part, f)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		params["input"] = "upload"
	} else {
		params["file"] = req.GetUri()
		params["input"] = "download"
	}

	for k, v := range params {
		if err := w.WriteField(k, v); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if err := w.Close(); err != nil {
		return nil, errors.WithStack(err)
	}

	httpReq, err := http.NewRequest("POST", "https://api.cloudconvert.com/convert", &b)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	httpReq.Header.Set("content-type", w.FormDataContentType())

	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	resCh := make(chan struct {
		r   *http.Response
		err error
	}, 1)

	go func() {
		res, err := client.Do(httpReq)
		resCh <- struct {
			r   *http.Response
			err error
		}{res, err}
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(httpReq)
		return nil, ctx.Err()
	case res := <-resCh:
		if res.err != nil {
			return nil, errors.WithStack(res.err)
		}
		if res.r.Body != nil {
			defer res.r.Body.Close()
		}

		if res.r.StatusCode != 200 {
			var data map[string]interface{}
			if err = json.NewDecoder(res.r.Body).Decode(&data); err != nil {
				return nil, errors.Wrap(err, "convert: cloudconvert conversion failed")
			}
			return nil, errors.Errorf("convert: cloudconvert conversion failed:\n%+v", data)
		}

		// Copy response to a new buffer as the HTTP response body will be closed
		var o bytes.Buffer
		if _, err := io.Copy(&o, res.r.Body); err != nil {
			return nil, errors.WithStack(err)
		}

		return &o, nil
	}
}

func (_ *CloudConvert) SupportedMimeTypes() []string {
	return []string{
		"application/msword",
		"application/rdf",
		"application/rdf+xml",
		"application/rtf",
		"application/xml",
		"application/x-mobipocket-ebook",
		"vnd.openxmlformats-officedocument",
		"x-iwork-pages",
		"image",
		"text",
	}
}

package process

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/fetcher"
	"github.com/arachnys/athenapdf/pkg/proto"
)

var (
	defaultConverter = "athenapdf"
	defaultMimeType  = "text/plain"
)

func Process(ctx context.Context, p *proto.Process) (io.Reader, error) {
	fr, mimeType, err := Fetch(ctx, p)
	if err != nil {
		return nil, err
	}

	if fr != nil {
		tmpFile, err := ioutil.TempFile("", "athenapdf-process")
		if err != nil {
			return nil, err
		}
		defer os.Remove(tmpFile.Name())

		if _, err := io.Copy(tmpFile, fr); err != nil {
			return nil, err
		}

		p.Conversion.Uri = fmt.Sprintf("file://%s", tmpFile.Name())
		if p.Conversion.GetMimeType() == "" {
			p.Conversion.MimeType = mimeType
		}
	}

	cr, err := Convert(ctx, p)
	if err != nil {
		return nil, err
	}

	return cr, nil
}

func Fetch(ctx context.Context, p *proto.Process) (io.Reader, string, error) {
	if p.Fetcher == "" {
		return nil, "", nil
	}

	f, err := fetcher.Get(p.Fetcher)
	if err != nil {
		return nil, "", err
	}

	return f.Fetch(ctx, p.Conversion.GetUri(), p.GetFetcherOptions())
}

func Convert(ctx context.Context, p *proto.Process) (io.Reader, error) {
	converterName := p.GetConverter()
	mimeType := p.Conversion.GetMimeType()

	// Attempt to use the specified converter if it supports the
	// specified mime type. Otherwise, attempt to use a converter
	// based on the mime type, if specified.
	var cvt converter.Converter
	var err error
	if converterName != "" {
		cvt, err = converter.Get(converterName)
		if err != nil {
			return nil, err
		}

		if mimeType != "" && !converter.IsMimeTypeSupported(cvt)(mimeType) {
			return nil, errors.Errorf("converter `%s` does not support mime type `%s`", converterName, mimeType)
		}
	} else if mimeType != "" {
		cvt, err = converter.GetFromMime(mimeType)
		if err != nil {
			return nil, err
		}
	} else {
		cvt, err = converter.Get(defaultConverter)
		if err != nil {
			return nil, err
		}
	}

	return cvt.Convert(ctx, p.Conversion, p.GetConverterOptions())
}

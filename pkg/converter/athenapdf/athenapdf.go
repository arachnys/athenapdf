package athenapdf

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"io"

	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/runner"
)

const (
	converterName = "athenapdf"
)

type AthenaPDF struct{}

func init() {
	converter.Register(converterName, &AthenaPDF{})
}

func (*AthenaPDF) String() string {
	return converterName
}

func (*AthenaPDF) Convert(ctx context.Context, req *proto.Conversion, opts map[string]*proto.Option) (io.Reader, error) {
	r := &runner.Runner{}
	exit, err := r.AutoTarget()
	defer func() {
		exit()
	}()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resCh := make(chan struct {
		r   io.Reader
		err error
	}, 1)

	go func() {
		b, err := r.Convert(req)
		resCh <- struct {
			r   io.Reader
			err error
		}{bytes.NewReader(b), err}
	}()

	select {
	case <-ctx.Done():
		// No clean up needed here as `exit` will be fired on return
		return nil, ctx.Err()
	case res := <-resCh:
		if res.err != nil {
			return nil, errors.WithStack(res.err)
		}
		return res.r, nil
	}
}

func (*AthenaPDF) SupportedMimeTypes() []string {
	return []string{
		"application/rdf",
		"application/rdf+xml",
		"application/rtf",
		"application/xml",
		"application/x-mimearchive",
		"multipart/related",
		"image",
		"text/css",
		"text/html",
	}
}

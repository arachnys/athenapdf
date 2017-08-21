package main

import (
	"container/heap"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"

	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/fetcher"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/uploader"
	"github.com/arachnys/athenapdf/pkg/uri"
)

type PDFService interface {
	Fetch(context.Context, *proto.Process) (io.Reader, string, error)
	Convert(context.Context, *proto.Process) (io.Reader, error)
	Upload(context.Context, *proto.Process, io.Reader) error
}

type pdfService struct{}

func (s pdfService) Fetch(ctx context.Context, p *proto.Process) (io.Reader, string, error) {
	f, err := fetcher.Get(p.GetFetcher())
	if err != nil {
		return nil, "", err
	}

	protocol, err := uri.Scheme(p.GetConversion().GetUri())
	if err != nil {
		return nil, "", err
	}

	if !fetcher.IsFetchable(f, protocol) {
		return nil, "", errors.Errorf(
			"target protocol `%s` is not supported",
			protocol,
		)
	}

	return f.Fetch(ctx, p.GetConversion().GetUri(), p.GetFetcherOptions())
}

func (s pdfService) Convert(ctx context.Context, p *proto.Process) (io.Reader, error) {
	converterName := p.GetConverter()

	if converterName == "" {
		// TODO: add priorities
		cq, err := converter.NewConverterQueue(p.GetConversion(), map[string]int{})
		if err != nil {
			return nil, err
		}

		attempts := 0
		for cq.Len() > 0 {
			item := heap.Pop(&cq).(*converter.ConverterQueueItem)
			r, err := item.Converter().Convert(ctx, p.GetConversion(), p.GetConverterOptions())
			attempts += 1

			converterName = "unknown"
			if s, ok := item.Converter().(fmt.Stringer); ok {
				converterName = s.String()
			}

			if err != nil {
				continue
			}

			p.Converter = converterName

			return r, nil
		}
	}

	c, err := converter.Get(converterName)
	if err != nil {
		return nil, err
	}

	if !converter.IsConvertable(c, p.GetConversion().GetMimeType()) {
		return nil, errors.Errorf(
			"mime type `%s` is not supported",
			p.GetConversion().GetMimeType(),
		)
	}

	return c.Convert(ctx, p.GetConversion(), p.GetConverterOptions())
}

func (s pdfService) Upload(ctx context.Context, p *proto.Process, cr io.Reader) error {
	u, err := uploader.Get(p.GetUploader())
	if err != nil {
		return err
	}

	return u.Upload(ctx, cr, p.GetUploaderOptions())
}

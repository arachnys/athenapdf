package main

import (
	"container/heap"
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/arachnys/athenapdf/pkg/converter"
	"github.com/arachnys/athenapdf/pkg/fetcher"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/uploader"
	"github.com/arachnys/athenapdf/pkg/uri"
)

type PDFService interface {
	Process(context.Context, *proto.Process) (io.Reader, bool, error)
}

type pdfService struct {
	tracer opentracing.Tracer
}

type CleanupFunc func() error

func (s pdfService) fetch(ctx context.Context, p *proto.Process) (CleanupFunc, error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		span := s.tracer.StartSpan(
			"fetch",
			opentracing.ChildOf(parentSpan.Context()),
		)
		span.LogFields(
			log.String("fetcher", p.GetFetcher()),
			log.String("uri", p.GetConversion().GetUri()),
		)
		defer span.Finish()
	}

	var cleanupFunc CleanupFunc = func() error { return nil }

	if p.GetFetcher() == "" {
		return cleanupFunc, nil
	}

	f, err := fetcher.Get(p.GetFetcher())
	if err != nil {
		return cleanupFunc, err
	}

	protocol, err := uri.Scheme(p.GetConversion().GetUri())
	if err != nil {
		return cleanupFunc, err
	}

	if !fetcher.IsFetchable(f, protocol) {
		return cleanupFunc, errors.Errorf(
			"target protocol `%s` is not supported",
			protocol,
		)
	}

	fr, mimeType, err := f.Fetch(ctx, p.GetConversion().GetUri(), p.GetFetcherOptions())
	if err != nil {
		return cleanupFunc, err
	}

	tmpFile, err := ioutil.TempFile("", "athenapdf-fetch")
	if err != nil {
		return cleanupFunc, errors.WithStack(err)
	}
	if tmpFile.Name() != "" {
		cleanupFunc = func() error {
			return os.Remove(tmpFile.Name())
		}
	}

	if _, err := io.Copy(tmpFile, fr); err != nil {
		return cleanupFunc, errors.WithStack(err)
	}

	// TODO: store the original URI
	p.Conversion.Uri = fmt.Sprintf("file://%s", tmpFile.Name())
	if p.GetConversion().GetMimeType() == "" {
		p.Conversion.MimeType = mimeType
	}

	return cleanupFunc, nil
}

func (s pdfService) convert(ctx context.Context, p *proto.Process) (io.Reader, error) {
	var span opentracing.Span
	converterName := p.GetConverter()

	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		span = s.tracer.StartSpan(
			"convert",
			opentracing.ChildOf(parentSpan.Context()),
		)
		defer func() {
			span.LogFields(
				log.String("mime type", p.GetConversion().GetMimeType()),
				log.String("uri", p.GetConversion().GetUri()),
			)
			span.Finish()
		}()
	}

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

			span.LogFields(log.String("converter", converterName))

			if err != nil {
				continue
			}

			span.LogFields(log.Int("attempts", attempts))
			return r, nil
		}
	}

	span.LogFields(log.String("converter", converterName))

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

func (s pdfService) upload(ctx context.Context, p *proto.Process, cr io.Reader) (bool, error) {
	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		span := s.tracer.StartSpan(
			"upload",
			opentracing.ChildOf(parentSpan.Context()),
		)
		span.LogFields(
			log.String("uploader", p.GetUploader()),
		)
		defer span.Finish()
	}

	if p.GetUploader() == "" {
		return false, nil
	}

	u, err := uploader.Get(p.GetUploader())
	if err != nil {
		return false, err
	}

	if err := u.Upload(ctx, cr, p.GetUploaderOptions()); err != nil {
		return false, err
	}

	return true, nil
}

func (s pdfService) Process(ctx context.Context, p *proto.Process) (io.Reader, bool, error) {
	var err error

	conversion := p.GetConversion()

	cleanupFunc, err := s.fetch(ctx, p)
	if err != nil {
		return nil, false, err
	}
	defer cleanupFunc()

	if conversion.GetMimeType() == "" {
		return nil, false, errors.New("unknown mime type")
	}

	cr, err := s.convert(ctx, p)
	if err != nil {
		return nil, false, err
	}

	uploaded, err := s.upload(ctx, p, cr)
	if err != nil {
		return nil, false, err
	}

	return cr, uploaded, nil
}

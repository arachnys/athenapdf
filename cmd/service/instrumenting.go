package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"io"

	"github.com/arachnys/athenapdf/pkg/proto"
)

type instrumentingMiddleware struct {
	tracer opentracing.Tracer
	next   PDFService
}

func (m instrumentingMiddleware) Fetch(ctx context.Context, p *proto.Process) (io.Reader, string, error) {
	var err error

	defer func() {
		if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
			span := m.tracer.StartSpan(
				"fetch",
				opentracing.ChildOf(parentSpan.Context()),
			)
			span.LogFields(log.String("uri", p.GetConversion().GetUri()))
			span.SetTag("fetcher", p.GetFetcher())
			if err != nil {
				span.LogFields(log.Error(err))
			}
			span.Finish()
		}
	}()

	r, mimeType, err := m.next.Fetch(ctx, p)
	return r, mimeType, err
}

func (m instrumentingMiddleware) Convert(ctx context.Context, p *proto.Process) (io.Reader, error) {
	var err error

	defer func() {
		if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
			span := m.tracer.StartSpan(
				"convert",
				opentracing.ChildOf(parentSpan.Context()),
			)
			span.LogFields(
				log.String("mime type", p.GetConversion().GetMimeType()),
				log.String("uri", p.GetConversion().GetUri()),
			)
			span.SetTag("converter", p.GetConverter())
			if err != nil {
				span.LogFields(log.Error(err))
			}
			span.Finish()
		}
	}()

	r, err := m.next.Convert(ctx, p)
	return r, err
}

func (m instrumentingMiddleware) Upload(ctx context.Context, p *proto.Process, r io.Reader) error {
	var err error

	defer func() {
		if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
			span := m.tracer.StartSpan(
				"upload",
				opentracing.ChildOf(parentSpan.Context()),
			)
			span.SetTag("uploader", p.GetUploader())
			if err != nil {
				span.LogFields(log.Error(err))
			}
			span.Finish()
		}
	}()

	err = m.next.Upload(ctx, p, r)
	return err
}

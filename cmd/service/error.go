package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"io"

	"github.com/arachnys/athenapdf/pkg/proto"
)

type errorMiddleware struct {
	tracer opentracing.Tracer
	next   PDFService
}

func (m errorMiddleware) Process(ctx context.Context, p *proto.Process) (io.Reader, bool, error) {
	var err error

	defer func() {
		if err != nil {
			if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
				parentSpan.LogFields(log.Error(err))
			}
		}
	}()

	r, u, err := m.next.Process(ctx, p)
	return r, u, err
}

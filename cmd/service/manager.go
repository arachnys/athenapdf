package main

import (
	"context"
	"io"

	"github.com/arachnys/athenapdf/pkg/processor"
	"github.com/arachnys/athenapdf/pkg/proto"
)

type managerMiddleware struct {
	manager processor.Manager
	next    PDFService
}

type convertProcess struct {
	svc PDFService
}

func (c convertProcess) Process(ctx context.Context, p *proto.Process) (io.Reader, error) {
	r, err := c.svc.Convert(ctx, p)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m managerMiddleware) Fetch(ctx context.Context, p *proto.Process) (io.Reader, string, error) {
	return m.next.Fetch(ctx, p)
}

func (m managerMiddleware) Convert(ctx context.Context, p *proto.Process) (io.Reader, error) {
	w, err := m.manager.Add(ctx, convertProcess{m.next}, p)
	if err != nil {
		return nil, err
	}

	select {
	case r := <-w.OutputReader:
		return r, nil
	case err := <-w.Err:
		return nil, err
	}
}

func (m managerMiddleware) Upload(ctx context.Context, p *proto.Process, r io.Reader) error {
	return m.next.Upload(ctx, p, r)
}

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

func (m managerMiddleware) Process(ctx context.Context, p *proto.Process) (io.Reader, bool, error) {
	w, err := m.manager.Add(ctx, m.next, p)
	if err != nil {
		return nil, false, err
	}

	select {
	case r := <-w.OutputReader:
		return r, false, nil
	case <-w.OutputUploaded:
		return nil, true, nil
	case err := <-w.Err:
		return nil, false, err
	}
}

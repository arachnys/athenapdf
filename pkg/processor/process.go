package processor

import (
	"context"
	"io"

	"github.com/arachnys/athenapdf/pkg/proto"
)

type Processor interface {
	Process(context.Context, *proto.Process) (io.Reader, error)
}

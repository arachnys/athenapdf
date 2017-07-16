package converter

import (
	"container/heap"
	"context"
	"github.com/pkg/errors"
	"io"
	"strings"
	"sync"

	"github.com/arachnys/athenapdf/pkg/proto"
)

var (
	converters   = make(map[string]Converter)
	convertersMu sync.RWMutex
)

type ConverterFunc func(ctx context.Context, req *proto.Conversion, opts map[string]*proto.Option) (io.Reader, error)

type Converter interface {
	Convert(context.Context, *proto.Conversion, map[string]*proto.Option) (io.Reader, error)
	SupportedMimeTypes() []string
}

func Register(converterName string, c Converter) error {
	convertersMu.Lock()
	defer convertersMu.Unlock()

	if converterName == "" {
		return ConverterError{err: errors.New("converter name is nil")}
	}
	if c == nil {
		return ConverterError{errors.New("converter is nil"), converterName}
	}

	converters[converterName] = c
	return nil
}

func List() map[string]Converter {
	return converters
}

func Get(converterName string) (Converter, error) {
	if c, ok := converters[converterName]; ok {
		return c, nil
	}
	return nil, ConverterError{
		err: errors.Errorf("converter `%s` does not exist", converterName),
	}
}

func NewConverterQueue(req *proto.Conversion, priorities map[string]int) (ConverterQueue, error) {
	cq := make(ConverterQueue, 0, len(List()))
	heap.Init(&cq)

	for name, c := range List() {
		if !IsConvertable(c, req.GetMimeType()) {
			continue
		}

		priority := 1
		if p, ok := priorities[name]; ok {
			priority = p
		}

		heap.Push(&cq, &ConverterQueueItem{
			converter: c,
			priority:  priority,
		})
	}

	if len(cq) == 0 {
		return nil, ConverterError{
			err: errors.Errorf("no converter with support for mime type `%s`", req.GetMimeType()),
		}
	}

	return cq, nil
}

func IsConvertable(c Converter, mimeType string) bool {
	for _, m := range c.SupportedMimeTypes() {
		if strings.Contains(strings.ToLower(mimeType), strings.ToLower(m)) {
			return true
		}
	}
	return false
}

package uploader

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"sync"

	"github.com/arachnys/athenapdf/pkg/proto"
)

var (
	uploaders   = make(map[string]Uploader)
	uploadersMu sync.RWMutex
)

type Uploader interface {
	Upload(context.Context, io.Reader, map[string]*proto.Option) error
}

func Register(uploaderName string, u Uploader) error {
	uploadersMu.Lock()
	defer uploadersMu.Unlock()

	if uploaderName == "" {
		return errors.New("uploader name is nil")
	}
	if u == nil {
		return errors.New("uploader is nil")
	}

	uploaders[uploaderName] = u
	return nil
}

func List() map[string]Uploader {
	return uploaders
}

func Get(uploaderName string) (Uploader, error) {
	if u, ok := uploaders[uploaderName]; ok {
		return u, nil
	}
	return nil, errors.Errorf("uploader `%s` does not exist", uploaderName)
}

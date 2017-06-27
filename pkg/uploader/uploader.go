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

type UploaderFunc func(context.Context, io.Reader, map[string]*proto.Option) error

type Uploader interface {
	Upload(context.Context, io.Reader, map[string]*proto.Option) error
}

func Register(uploaderName string, u Uploader) error {
	uploadersMu.Lock()
	defer uploadersMu.Unlock()

	if uploaderName == "" {
		return UploaderError{err: errors.New("uploader name is nil")}
	}
	if u == nil {
		return UploaderError{errors.New("uploader is nil"), uploaderName}
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
	return nil, UploaderError{err: errors.Errorf("uploader `%s` does not exist", uploaderName)}
}

func Upload(uploaderName string) UploaderFunc {
	return func(ctx context.Context, r io.Reader, opts map[string]*proto.Option) error {
		u, err := Get(uploaderName)
		if err != nil {
			return err
		}

		return u.Upload(ctx, r, opts)
	}
}

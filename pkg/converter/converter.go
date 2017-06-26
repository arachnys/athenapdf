package converter

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"mime"
	"strings"
	"sync"

	"github.com/arachnys/athenapdf/pkg/proto"
)

var (
	converters   = make(map[string]Converter)
	convertersMu sync.RWMutex
)

type Converter interface {
	Convert(context.Context, *proto.Conversion, map[string]*proto.Option) (io.Reader, error)
	SupportedMimeTypes() []string
}

func Register(converterName string, c Converter) error {
	convertersMu.Lock()
	defer convertersMu.Unlock()

	if converterName == "" {
		return errors.New("converter name is nil")
	}
	if c == nil {
		return errors.New("converter is nil")
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
	return nil, errors.Errorf("converter `%s` does not exist", converterName)
}

func GetFromMime(targetMimeType string) (Converter, error) {
	for _, c := range List() {
		if IsMimeTypeSupported(c)(targetMimeType) {
			return c, nil
		}
	}
	return nil, errors.Errorf("no converter with support for mime type `%s`", targetMimeType)
}

func IsMimeTypeSupported(c Converter) func(string) bool {
	return func(targetMimeType string) bool {
		for _, mimeType := range c.SupportedMimeTypes() {
			if strings.Contains(strings.ToLower(targetMimeType), strings.ToLower(mimeType)) {
				return true
			}
		}
		return false
	}
}

func IsLocalConversion(uri string) bool {
	if strings.Index(strings.ToLower(uri), "file://") == 0 {
		return true
	}
	return false
}

func ExtensionByMimeType(mimeType string) string {
	e, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(e) == 0 {
		return ""
	}
	return strings.TrimPrefix(e[0], ".")
}

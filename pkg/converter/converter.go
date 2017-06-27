package converter

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"strings"
	"sync"

	"github.com/arachnys/athenapdf/pkg/mime"
	"github.com/arachnys/athenapdf/pkg/proto"
)

const (
	defaultMimeType = "text/plain; charset=\"UTF-8\""
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
		return errors.New("register: converter name is nil")
	}
	if c == nil {
		return errors.New("register: converter is nil")
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
	return nil, errors.Errorf("get: converter `%s` does not exist", converterName)
}

func GetFromConversion(req *proto.Conversion) (Converter, error) {
	for _, c := range List() {
		if IsConvertable(c, req.GetMimeType()) {
			return c, nil
		}
	}
	return nil, errors.Errorf("getfromconversion: no converter with support for mime type `%s`", req.GetMimeType())
}

func GetFromConversionExcluding(req *proto.Conversion) func([]string) (Converter, error) {
	return func(excludedConverters []string) (Converter, error) {
		for converterName, c := range List() {
			for _, excludedConverter := range excludedConverters {
				if converterName == excludedConverter {
					continue
				}

				if IsConvertable(c, req.GetMimeType()) {
					return c, nil
				}
			}
		}

		return nil, errors.Errorf(
			"getfromconversionexcluding: no converter left with support for mime type `%s`",
			req.GetMimeType(),
		)
	}
}

func Convert(converterName string) func(context.Context, *proto.Conversion, map[string]*proto.Option) (io.Reader, error) {
	return func(ctx context.Context, req *proto.Conversion, opts map[string]*proto.Option) (io.Reader, error) {
		var c Converter
		var err error

		if req.GetMimeType() == "" {
			if !IsLocal(req) {
				return nil, errors.Errorf(
					"convert: %s: mime type must be provided for non-local conversions",
					converterName,
				)
			}

			mimeType, err := mime.TypeFromFile(req.GetUri())
			if err != nil {
				return nil, err
			}
			if mimeType == "" {
				mimeType = defaultMimeType
			}

			req.MimeType = mimeType
		}

		if converterName != "" {
			c, err = Get(converterName)
			if err != nil {
				return nil, err
			}

			if !IsConvertable(c, req.GetMimeType()) {
				return nil, errors.Errorf(
					"convert: %s: mime type `%s` is not supported",
					converterName,
					req.GetMimeType(),
				)
			}
		}

		if c == nil {
			if c, err = GetFromConversion(req); err != nil {
				return nil, err
			}
		}

		return c.Convert(ctx, req, opts)
	}
}

func IsConvertable(c Converter, mimeType string) bool {
	for _, m := range c.SupportedMimeTypes() {
		if strings.Contains(strings.ToLower(mimeType), strings.ToLower(m)) {
			return true
		}
	}
	return false
}

func IsLocal(req *proto.Conversion) bool {
	if strings.Index(strings.ToLower(req.GetUri()), "file://") == 0 {
		return true
	}
	return false
}

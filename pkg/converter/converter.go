package converter

import (
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

func GetFromConversion(req *proto.Conversion) (Converter, error) {
	for _, c := range List() {
		if IsConvertable(c, req.GetMimeType()) {
			return c, nil
		}
	}
	return nil, ConverterError{
		err: errors.Errorf("no converter with support for mime type `%s`", req.GetMimeType()),
	}
}

func IsExcluded(converterName string, excludedConverters []string) bool {
	for _, excludedConverter := range excludedConverters {
		if converterName == excludedConverter {
			return true
		}
	}
	return false
}

func GetFromConversionExcluding(req *proto.Conversion) func([]string) (Converter, error) {
	return func(excludedConverters []string) (Converter, error) {
		for converterName, c := range List() {
			if IsConvertable(c, req.GetMimeType()) && !IsExcluded(converterName, excludedConverters) {
				return c, nil
			}
		}

		return nil, ConverterError{
			err: errors.Errorf(
				"no converter left with support for mime type `%s`",
				req.GetMimeType(),
			),
		}
	}
}

func Convert(converterName string) ConverterFunc {
	return func(ctx context.Context, req *proto.Conversion, opts map[string]*proto.Option) (io.Reader, error) {
		c, err := Get(converterName)
		if err != nil {
			return nil, err
		}

		if !IsConvertable(c, req.GetMimeType()) {
			return nil, ConverterError{
				errors.Errorf(
					"mime type `%s` is not supported",
					req.GetMimeType(),
				),
				converterName,
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

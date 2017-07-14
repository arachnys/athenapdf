package converter

import (
	"context"
	"io"
	"testing"

	"github.com/arachnys/athenapdf/pkg/proto"
)

type MockConverter struct{}

func (*MockConverter) Convert(_ context.Context, _ *proto.Conversion, _ map[string]*proto.Option) (io.Reader, error) {
	return nil, nil
}

func (*MockConverter) SupportedMimeTypes() []string {
	return []string{}
}

func TestRegister(t *testing.T) {
	c := &MockConverter{}

	testCases := []struct {
		name          string
		converterName string
		converter     Converter
	}{
		{"no converter name", "", c},
		{"no converter", "mock", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Register(tc.converterName, tc.converter) == nil {
				t.Errorf("expected error registering converter without correct arguments")
			}
		})
	}

	t.Run("converter with name", func(t *testing.T) {
		converterName := "mock"

		if err := Register(converterName, c); err != nil {
			t.Fatalf("failed to register converter, unexpected error: %+v", err)
		}

		got, err := Get(converterName)
		if err != nil {
			t.Errorf("failed to retrieve registered converter, unexpected error: %+v", err)
		}

		if got != c {
			t.Errorf("got %+v; want %+v", got, c)
		}
	})
}

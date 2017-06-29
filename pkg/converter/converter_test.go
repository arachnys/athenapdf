package converter

import (
	"context"
	"io"
	"testing"

	"github.com/arachnys/athenapdf/pkg/proto"
)

type MockConverter struct{}

func (_ *MockConverter) Convert(_ context.Context, _ *proto.Conversion, _ map[string]*proto.Option) (io.Reader, error) {
	return nil, nil
}

func (_ *MockConverter) SupportedMimeTypes() []string {
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

func TestIsLocal(t *testing.T) {
	testCases := []struct {
		conversion *proto.Conversion
		want       bool
	}{
		{&proto.Conversion{Uri: "https://www.athenapdf.com/"}, false},
		{&proto.Conversion{Uri: "ftp://fyianlai.com"}, false},
		{&proto.Conversion{Uri: "s3://s3.amazonaws.com/"}, false},
		{&proto.Conversion{Uri: "file://home/athena/index.html"}, true},
		{&proto.Conversion{Uri: "file:///home/athena/index.html"}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.conversion.GetUri(), func(t *testing.T) {
			if got := IsLocal(tc.conversion); got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}
}

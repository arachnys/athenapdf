package uploader

import (
	"context"
	"io"
	"testing"

	"github.com/arachnys/athenapdf/pkg/proto"
)

type MockUploader struct{}

func (*MockUploader) Upload(_ context.Context, _ io.Reader, _ map[string]*proto.Option) error {
	return nil
}

func TestRegister(t *testing.T) {
	u := &MockUploader{}

	testCases := []struct {
		name         string
		uploaderName string
		uploader     Uploader
	}{
		{"no uploader name", "", u},
		{"no uploader", "mock", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Register(tc.uploaderName, tc.uploader) == nil {
				t.Errorf("expected error registering uploader without correct arguments")
			}
		})
	}

	t.Run("uploader with name", func(t *testing.T) {
		uploaderName := "mock"

		if err := Register(uploaderName, u); err != nil {
			t.Fatalf("failed to register uploader, unexpected error: %+v", err)
		}

		got, err := Get(uploaderName)
		if err != nil {
			t.Errorf("failed to retrieve registered uploader, unexpected error: %+v", err)
		}

		if got != u {
			t.Errorf("got %+v; want %+v", got, u)
		}
	})
}

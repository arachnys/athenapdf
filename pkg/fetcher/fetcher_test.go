package fetcher

import (
	"context"
	"io"
	"testing"

	"github.com/arachnys/athenapdf/pkg/proto"
)

type MockFetcher struct{}

func (_ *MockFetcher) Fetch(_ context.Context, _ string, _ map[string]*proto.Option) (io.Reader, string, error) {
	return nil, "", nil
}

func (_ *MockFetcher) SupportedProtocols() []string {
	return []string{}
}

func TestRegister(t *testing.T) {
	f := &MockFetcher{}

	testCases := []struct {
		name        string
		fetcherName string
		fetcher     Fetcher
	}{
		{"no fetcher name", "", f},
		{"no fetcher", "mock", nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if Register(tc.fetcherName, tc.fetcher) == nil {
				t.Errorf("expected error registering fetcher without correct arguments")
			}
		})
	}

	t.Run("fetcher with name", func(t *testing.T) {
		fetcherName := "mock"

		if err := Register(fetcherName, f); err != nil {
			t.Fatalf("failed to register fetcher, unexpected error: %+v", err)
		}

		got, err := Get(fetcherName)
		if err != nil {
			t.Errorf("failed to retrieve registered fetcher, unexpected error: %+v", err)
		}

		if got != f {
			t.Errorf("got %+v; want %+v", got, f)
		}
	})
}

func TestToProtocol(t *testing.T) {
	testCases := []struct {
		url  string
		want string
	}{
		{"http://www.arachnys.com", "http"},
		{"https://www.athenapdf.com/", "https"},
		{"ftp://fyianlai.com", "ftp"},
		{"s3://s3.amazonaws.com/", "s3"},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			got, err := ToProtocol(tc.url)
			if err != nil {
				t.Fatalf("failed to get protocol from url, unexpected error: %+v", err)
			}

			if got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}

	t.Run("invalid url", func(t *testing.T) {
		_, err := ToProtocol("% invalid url")
		if err == nil {
			t.Fatalf("expected error with invalid url")
		}
	})
}

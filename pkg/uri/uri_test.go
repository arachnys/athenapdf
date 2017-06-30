package uri

import (
	"testing"
)

func TestIsLocal(t *testing.T) {
	testCases := []struct {
		uri  string
		want bool
	}{
		{"https://www.athenapdf.com/", false},
		{"ftp://fyianlai.com", false},
		{"s3://s3.amazonaws.com/", false},
		{"file://home/athena/index.html", true},
		{"file:///home/athena/index.html", true},
	}

	for _, tc := range testCases {
		t.Run(tc.uri, func(t *testing.T) {
			if got := IsLocal(tc.uri); got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}
}

func TestScheme(t *testing.T) {
	testCases := []struct {
		url  string
		want string
	}{
		{"file:///home/athena/index.html", "file"},
		{"http://www.arachnys.com", "http"},
		{"https://www.athenapdf.com/", "https"},
		{"ftp://fyianlai.com", "ftp"},
		{"s3://s3.amazonaws.com/", "s3"},
	}

	for _, tc := range testCases {
		t.Run(tc.url, func(t *testing.T) {
			got, err := Scheme(tc.url)
			if err != nil {
				t.Fatalf("failed to get scheme from url, unexpected error: %+v", err)
			}

			if got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}

	t.Run("invalid url", func(t *testing.T) {
		_, err := Scheme("% invalid url")
		if err == nil {
			t.Fatalf("expected error with invalid url")
		}
	})
}

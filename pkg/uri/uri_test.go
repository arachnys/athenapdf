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

func TestToLocal(t *testing.T) {
	testCases := []struct {
		uri  string
		want string
	}{
		{"relative/path/to/local/file", "file://relative/path/to/local/file"},
		{"/absolute/path/to/local/file", "file:///absolute/path/to/local/file"},
		{"file:///path/to/local/file", "file:///path/to/local/file"},
		{"ftp://fyianlai.com", "file://fyianlai.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.uri, func(t *testing.T) {
			got, err := ToLocal(tc.uri)
			if err != nil {
				t.Fatalf("failed to get convert uri scheme to local file, unexpected error: %+v", err)
			}

			if got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}

	t.Run("invalid uri", func(t *testing.T) {
		_, err := ToLocal("% invalid uri")
		if err == nil {
			t.Fatalf("expected error with invalid uri")
		}
	})
}

func TestScheme(t *testing.T) {
	testCases := []struct {
		uri  string
		want string
	}{
		{"file:///home/athena/index.html", "file"},
		{"http://www.arachnys.com", "http"},
		{"https://www.athenapdf.com/", "https"},
		{"ftp://fyianlai.com", "ftp"},
		{"s3://s3.amazonaws.com/", "s3"},
	}

	for _, tc := range testCases {
		t.Run(tc.uri, func(t *testing.T) {
			got, err := Scheme(tc.uri)
			if err != nil {
				t.Fatalf("failed to get scheme from uri, unexpected error: %+v", err)
			}

			if got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}

	t.Run("invalid uri", func(t *testing.T) {
		_, err := Scheme("% invalid uri")
		if err == nil {
			t.Fatalf("expected error with invalid uri")
		}
	})
}

func TestRemoveScheme(t *testing.T) {
	testCases := []struct {
		uri  string
		want string
	}{
		{"file:///home/athena/index.html", "/home/athena/index.html"},
		{"http://www.arachnys.com", "//www.arachnys.com"},
		{"https://www.athenapdf.com/", "//www.athenapdf.com/"},
		{"ftp://fyianlai.com", "//fyianlai.com"},
		{"s3://s3.amazonaws.com/", "//s3.amazonaws.com/"},
	}

	for _, tc := range testCases {
		t.Run(tc.uri, func(t *testing.T) {
			got, err := RemoveScheme(tc.uri)
			if err != nil {
				t.Fatalf("failed to remove scheme from uri, unexpected error: %+v", err)
			}

			if got != tc.want {
				t.Errorf("got %+v; want %+v", got, tc.want)
			}
		})
	}

	t.Run("invalid uri", func(t *testing.T) {
		_, err := RemoveScheme("% invalid uri")
		if err == nil {
			t.Fatalf("expected error with invalid uri")
		}
	})
}

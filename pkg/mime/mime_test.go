package mime

import (
	"testing"
)

func TestToExtension(t *testing.T) {
	testCases := []struct {
		mimeType string
		want     []string
	}{
		{"application/msword", []string{"doc"}},
		{"application/pdf", []string{"pdf"}},
		{"text/csv", []string{"csv"}},
		{"text/html", []string{"htm", "html"}},
		{"text/plain", []string{"txt"}},
		{"application/vnd.oasis.opendocument.text", []string{"odt"}},
	}

	for _, tc := range testCases {
		t.Run(tc.mimeType, func(t *testing.T) {
			got := ToExtension(tc.mimeType)

			for _, want := range tc.want {
				if got == want {
					return
				}
			}

			t.Errorf("got %+v; want %+v", got, tc.want)
		})
	}
}

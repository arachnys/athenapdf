package mime

import (
	"github.com/pkg/errors"
	"io"
	stdmime "mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultMimeType = "text/plain; charset=\"UTF-8\""
)

func ToExtension(mimeType string) string {
	e, err := stdmime.ExtensionsByType(mimeType)
	if err != nil || len(e) == 0 {
		return ""
	}
	return strings.TrimPrefix(e[0], ".")
}

func TypeFromReader(r io.Reader) (string, error) {
	mimeTypeBuf := make([]byte, 0, 512)
	n, err := r.Read(mimeTypeBuf)
	if err != nil {
		return "", errors.WithStack(err)
	}
	mimeType := defaultMimeType
	if n >= 512 {
		mimeType = http.DetectContentType(mimeTypeBuf)
	}
	return mimeType, nil
}

func TypeFromFile(filePath string) (string, error) {
	if mimeType := stdmime.TypeByExtension(filepath.Ext(filePath)); mimeType != "" {
		return mimeType, nil
	}

	f, err := os.Open(filePath)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if f != nil {
		defer f.Close()
	}

	mimeType, err := TypeFromReader(f)
	if err != nil {
		return "", err
	}

	return mimeType, nil
}

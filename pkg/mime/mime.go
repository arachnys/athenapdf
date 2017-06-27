package mime

import (
	"bytes"
	"io"
	stdmime "mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ToExtension(mimeType string) string {
	e, err := stdmime.ExtensionsByType(mimeType)
	if err != nil || len(e) == 0 {
		return ""
	}
	return strings.TrimPrefix(e[0], ".")
}

func TypeFromReader(r io.Reader) (string, error) {
	contentTypeBuf := bytes.NewBuffer(make([]byte, 0, 512))
	if _, err := io.Copy(contentTypeBuf, r); err != nil {
		return "", err
	}
	return http.DetectContentType(contentTypeBuf.Bytes()), nil
}

func TypeFromFile(filePath string) (string, error) {
	if mimeType := stdmime.TypeByExtension(filepath.Ext(filePath)); mimeType != "" {
		return mimeType, nil
	}

	f, err := os.Open(filePath)
	if err != nil {
		return "", err
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

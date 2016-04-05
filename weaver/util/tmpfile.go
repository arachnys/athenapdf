package util

import (
	"io"
	"io/ioutil"
	"path/filepath"
)

// NewTmpFile returns the absolute path to a temporary file containing bytes
// copied from an io.Reader.
// It will return an empty string if there is an error.
func NewTmpFile(body io.Reader) (string, error) {
	file, err := ioutil.TempFile("/tmp", "tmp")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return "", err
	}

	path, err := filepath.Abs(file.Name())
	if err != nil {
		return "", err
	}

	return path, nil
}

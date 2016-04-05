package util

import (
	"log"
	"net/http"
)

// HandleOctetStream returns the absolute path to a temporary file containing
// a downloaded HTTP resource if the URL provided has a content-type of
// "application/octet-stream" (i.e. its body contains arbitary binary data).
// Otherwise, it will return an empty string.
// An empty string is always returned when there is an error.
func HandleOctetStream(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.Header.Get("Content-Type") == "application/octet-stream" {
		log.Println("application/octet-stream detected, saving, and converting locally")
		p, err := NewTmpFile(res.Body)
		if err != nil {
			return "", err
		}
		return p, nil
	}

	return "", nil
}

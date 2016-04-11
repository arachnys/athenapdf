package util

import (
	"golang.org/x/net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
)

// HandleOctetStream returns the absolute path to a temporary file containing
// a downloaded HTTP resource if the URL provided has a content-type of
// "application/octet-stream" (i.e. its body contains arbitary binary data).
// Otherwise, it will return an empty string.
// An empty string is always returned when there is an error.
func HandleOctetStream(url string) (string, error) {
	opts := cookiejar.Options{PublicSuffixList: publicsuffix.List}
	jar, err := cookiejar.New(&opts)
	if err != nil {
		return "", err
	}
	client := http.Client{Jar: jar}
	res, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if res != nil {
		defer res.Body.Close()
	}

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

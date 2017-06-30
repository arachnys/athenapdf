package uri

import (
	"net/url"
)

func IsLocal(uri string) bool {
	if scheme, _ := Scheme(uri); scheme == "file" {
		return true
	}
	return false
}

func Scheme(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return u.Scheme, nil
}

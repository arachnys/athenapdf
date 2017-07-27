package uri

import (
	"github.com/pkg/errors"
	"net/url"
)

func IsLocal(uri string) bool {
	if scheme, _ := Scheme(uri); scheme == "file" {
		return true
	}
	return false
}

func ToLocal(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", errors.WithStack(err)
	}
	u.Scheme = "file"
	return u.String(), nil
}

func Scheme(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return u.Scheme, nil
}

func RemoveScheme(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", errors.WithStack(err)
	}
	u.Scheme = ""
	return u.String(), nil
}

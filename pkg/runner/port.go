package runner

import (
	"github.com/pkg/errors"
	"net"
)

func getRandomPort() (string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer l.Close()

	_, port, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return "", errors.WithStack(err)
	}

	return port, nil
}

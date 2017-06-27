package fetcher

import (
	"fmt"
)

type FetcherError struct {
	err         error
	fetcherName string
}

func (e FetcherError) Error() string {
	if e.fetcherName != "" {
		return fmt.Sprintf("%s: %s", e.fetcherName, e.err.Error())
	}
	return e.err.Error()
}

func (e FetcherError) Cause() error {
	return e.err
}

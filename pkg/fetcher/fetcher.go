package fetcher

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"strings"
	"sync"

	"github.com/arachnys/athenapdf/pkg/proto"
)

var (
	fetchers   = make(map[string]Fetcher)
	fetchersMu sync.RWMutex
)

type FetcherFunc func(context.Context, string, map[string]*proto.Option) (io.Reader, string, error)

type Fetcher interface {
	Fetch(context.Context, string, map[string]*proto.Option) (io.Reader, string, error)
	SupportedProtocols() []string
}

func Register(fetcherName string, f Fetcher) error {
	fetchersMu.Lock()
	defer fetchersMu.Unlock()
	if fetcherName == "" {
		return FetcherError{err: errors.New("fetcher name is nil")}
	}
	if f == nil {
		return FetcherError{err: errors.New("fetcher is nil")}
	}
	fetchers[fetcherName] = f
	return nil
}

func List() map[string]Fetcher {
	return fetchers
}

func Get(fetcherName string) (Fetcher, error) {
	if f, ok := fetchers[fetcherName]; ok {
		return f, nil
	}
	return nil, FetcherError{err: errors.Errorf("fetcher `%s` does not exist", fetcherName)}
}

func IsFetchable(f Fetcher, protocol string) bool {
	for _, p := range f.SupportedProtocols() {
		if strings.HasPrefix(strings.ToLower(protocol), strings.ToLower(p)) {
			return true
		}
	}
	return false
}

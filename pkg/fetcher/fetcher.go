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

type Fetcher interface {
	Fetch(context.Context, string, map[string]*proto.Option) (io.Reader, string, error)
	SupportedProtocols() []string
}

func Register(fetcherName string, f Fetcher) error {
	fetchersMu.Lock()
	defer fetchersMu.Unlock()

	if fetcherName == "" {
		return errors.New("fetcher name is nil")
	}
	if f == nil {
		return errors.New("fetcher is nil")
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
	return nil, errors.Errorf("fetcher `%s` does not exist", fetcherName)
}

func IsURIProtocolSupported(f Fetcher) func(string) bool {
	return func(uri string) bool {
		for _, protocol := range f.SupportedProtocols() {
			if strings.HasPrefix(strings.ToLower(uri), strings.ToLower(protocol)) {
				return true
			}
		}
		return false
	}
}

package fetcher

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"net/url"
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

func Fetch(fetcherName string) FetcherFunc {
	return func(ctx context.Context, target string, opts map[string]*proto.Option) (io.Reader, string, error) {
		f, err := Get(fetcherName)
		if err != nil {
			return nil, "", err
		}

		protocol, err := ToProtocol(target)
		if err != nil {
			return nil, "", err
		}

		if !IsFetchable(f, protocol) {
			return nil, "", FetcherError{
				errors.Errorf(
					"target protocol `%s` is not supported",
					protocol,
				),
				fetcherName,
			}
		}

		return f.Fetch(ctx, target, opts)
	}
}

func IsFetchable(f Fetcher, protocol string) bool {
	for _, p := range f.SupportedProtocols() {
		if strings.HasPrefix(strings.ToLower(protocol), strings.ToLower(p)) {
			return true
		}
	}
	return false
}

func ToProtocol(target string) (string, error) {
	u, err := url.Parse(target)
	if err != nil {
		return "", err
	}

	return u.Scheme, nil
}

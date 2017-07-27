package tally

import (
	"github.com/uber-go/tally"

	"github.com/uber/jaeger-lib/metrics"
)

// Wrap takes a tally Scope and returns jaeger-lib metrics.Factory.
func Wrap(scope tally.Scope) metrics.Factory {
	return &factory{
		tally: scope,
	}
}

// TODO implement support for tags if tally.Scope does not support them
type factory struct {
	tally tally.Scope
}

func (f *factory) Counter(name string, tags map[string]string) metrics.Counter {
	scope := f.tally
	if len(tags) > 0 {
		scope = scope.Tagged(tags)
	}
	return NewCounter(scope.Counter(name))
}

func (f *factory) Gauge(name string, tags map[string]string) metrics.Gauge {
	scope := f.tally
	if len(tags) > 0 {
		scope = scope.Tagged(tags)
	}
	return NewGauge(scope.Gauge(name))
}

func (f *factory) Timer(name string, tags map[string]string) metrics.Timer {
	scope := f.tally
	if len(tags) > 0 {
		scope = scope.Tagged(tags)
	}
	return NewTimer(scope.Timer(name))
}

func (f *factory) Namespace(name string, tags map[string]string) metrics.Factory {
	return &factory{
		tally: f.tally.SubScope(name).Tagged(tags),
	}
}

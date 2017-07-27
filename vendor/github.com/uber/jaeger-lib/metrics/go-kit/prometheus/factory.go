// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package prometheus

import (
	"strings"

	"github.com/go-kit/kit/metrics"
	kitprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/uber/jaeger-lib/metrics/go-kit"
)

var normalizer = strings.NewReplacer(
	".", "_",
	"-", "_",
)

// NewFactory creates a new metrics factory using go-kit prometheus package.
// buckets define the buckets into which histogram observations are counted.
// If buckets == nil, the default value prometheus.DefBuckets is used.
func NewFactory(namespace, subsystem string, buckets []float64) xkit.Factory {
	return &factory{
		namespace: namespace,
		subsystem: subsystem,
		buckets:   buckets,
	}
}

type factory struct {
	namespace string
	subsystem string
	buckets   []float64
}

func (f *factory) Counter(name string) metrics.Counter {
	opts := prometheus.CounterOpts{
		Namespace: f.namespace,
		Subsystem: f.subsystem,
		Name:      normalizer.Replace(name),
		Help:      name,
	}
	return kitprom.NewCounterFrom(opts, nil)
}

func (f *factory) Histogram(name string) metrics.Histogram {
	opts := prometheus.HistogramOpts{
		Namespace: f.namespace,
		Subsystem: f.subsystem,
		Name:      normalizer.Replace(name),
		Help:      name,
		Buckets:   f.buckets,
	}
	return kitprom.NewHistogramFrom(opts, nil)
}

func (f *factory) Gauge(name string) metrics.Gauge {
	opts := prometheus.GaugeOpts{
		Namespace: f.namespace,
		Subsystem: f.subsystem,
		Name:      normalizer.Replace(name),
		Help:      name,
	}
	return kitprom.NewGaugeFrom(opts, nil)
}

func (f *factory) Capabilities() xkit.Capabilities {
	return xkit.Capabilities{Tagging: true}
}

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

package expvar

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/expvar"

	"github.com/uber/jaeger-lib/metrics/go-kit"
)

// NewFactory creates a new metrics factory using go-kit expvar package.
// buckets is the number of buckets to be used in histograms.
func NewFactory(buckets int) xkit.Factory {
	return factory{
		buckets: buckets,
	}
}

type factory struct {
	buckets int
}

func (f factory) Counter(name string) metrics.Counter {
	return expvar.NewCounter(name)
}

func (f factory) Histogram(name string) metrics.Histogram {
	return expvar.NewHistogram(name, f.buckets)
}

func (f factory) Gauge(name string) metrics.Gauge {
	return expvar.NewGauge(name)
}

func (f factory) Capabilities() xkit.Capabilities {
	return xkit.Capabilities{Tagging: false}
}

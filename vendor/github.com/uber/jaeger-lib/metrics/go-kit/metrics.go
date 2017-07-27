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

package xkit

import (
	"time"

	kit "github.com/go-kit/kit/metrics"
)

// Counter is an adapter from go-kit Counter to jaeger-lib Counter
type Counter struct {
	counter kit.Counter
}

// NewCounter creates a new Counter
func NewCounter(counter kit.Counter) *Counter {
	return &Counter{counter: counter}
}

// Inc adds the given value to the counter.
func (c *Counter) Inc(delta int64) {
	c.counter.Add(float64(delta))
}

// Gauge is an adapter from go-kit Gauge to jaeger-lib Gauge
type Gauge struct {
	gauge kit.Gauge
}

// NewGauge creates a new Gauge
func NewGauge(gauge kit.Gauge) *Gauge {
	return &Gauge{gauge: gauge}
}

// Update the gauge to the value passed in.
func (g *Gauge) Update(value int64) {
	g.gauge.Set(float64(value))
}

// Timer is an adapter from go-kit Histogram to jaeger-lib Timer
type Timer struct {
	hist kit.Histogram
}

// NewTimer creates a new Timer
func NewTimer(hist kit.Histogram) *Timer {
	return &Timer{hist: hist}
}

// Record saves the time passed in.
func (t *Timer) Record(delta time.Duration) {
	t.hist.Observe(delta.Seconds())
}

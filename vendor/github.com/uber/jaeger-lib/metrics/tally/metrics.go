package tally

import (
	"time"

	"github.com/uber-go/tally"
)

// Counter is an adapter from go-tally Counter to jaeger-lib Counter
type Counter struct {
	counter tally.Counter
}

// NewCounter creates a new Counter
func NewCounter(counter tally.Counter) *Counter {
	return &Counter{counter: counter}
}

// Inc adds the given value to the counter.
func (c *Counter) Inc(delta int64) {
	c.counter.Inc(delta)
}

// Gauge is an adapter from go-tally Gauge to jaeger-lib Gauge
type Gauge struct {
	gauge tally.Gauge
}

// NewGauge creates a new Gauge
func NewGauge(gauge tally.Gauge) *Gauge {
	return &Gauge{gauge: gauge}
}

// Update the gauge to the value passed in.
func (g *Gauge) Update(value int64) {
	g.gauge.Update(float64(value))
}

// Timer is an adapter from go-tally Histogram to jaeger-lib Timer
type Timer struct {
	timer tally.Timer
}

// NewTimer creates a new Timer
func NewTimer(timer tally.Timer) *Timer {
	return &Timer{timer: timer}
}

// Record saves the time passed in.
func (t *Timer) Record(delta time.Duration) {
	t.timer.Record(delta)
}

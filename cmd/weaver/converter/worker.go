package converter

import (
	"errors"
	"log"
	"time"
)

var (
	ErrConversionTimeout = errors.New("conversion timed out")
)

type Worker struct {
	id int
}

func InitWorkers(maxWorkers, maxQueue, t int) chan<- Work {
	wq := make(chan Work, maxQueue)

	for i := 0; i < maxWorkers; i++ {
		w := Worker{i}
		go func(wq <-chan Work, w Worker, t int) {
			for work := range wq {
				log.Printf("[Worker #%d] processing conversion job (pending conversions: %d)\n", w.id, len(wq))
				work.Process(t)
			}
		}(wq, w, t)
	}

	return wq
}

type Work struct {
	converter Converter
	source    ConversionSource
	out       chan []byte
	err       chan error
	uploaded  chan struct{}
	done      chan struct{}
}

func NewWork(wq chan<- Work, c Converter, s ConversionSource) Work {
	w := Work{}
	w.converter = c
	w.source = s
	w.out = make(chan []byte, 1)
	w.err = make(chan error, 1)
	w.uploaded = make(chan struct{}, 1)
	w.done = make(chan struct{}, 1)
	go func(wq chan<- Work, w Work) {
		wq <- w
	}(wq, w)
	return w
}

func (w Work) Process(t int) {
	done := make(chan struct{}, 1)
	defer close(done)

	wout := make(chan []byte, 1)
	werr := make(chan error, 1)

	go func(w Work, done <-chan struct{}, wout chan<- []byte, werr chan<- error) {
		out, err := w.converter.Convert(w.source, done)
		if err != nil {
			werr <- err
			return
		}

		uploaded, err := w.converter.Upload(out)
		if err != nil {
			werr <- err
			return
		}

		if uploaded {
			close(w.uploaded)
			return
		}

		wout <- out
	}(w, done, wout, werr)

	select {
	case <-w.Cancelled():
	case <-w.Uploaded():
	case out := <-wout:
		w.out <- out
	case err := <-werr:
		w.err <- err
	case <-time.After(time.Second * time.Duration(t)):
		w.err <- ErrConversionTimeout
	}
}

// Success returns a channel that will be used for publishing the output of a
// conversion.
func (w Work) Success() <-chan []byte {
	return w.out
}

// Error returns a channel that will be used for publishing errors from a
// conversion.
func (w Work) Error() <-chan error {
	return w.err
}

// Uploaded returns a channel that will be closed when a conversion has been
// uploaded.
func (w Work) Uploaded() <-chan struct{} {
	return w.uploaded
}

// Cancel will close the done channel. This will indicate to child Goroutines
// that the job has been terminated, and the results are no longer needed.
func (w Work) Cancel() {
	close(w.done)
}

// Cancelled returns a channel that will indicate when a job has been completed.
func (w Work) Cancelled() <-chan struct{} {
	return w.done
}

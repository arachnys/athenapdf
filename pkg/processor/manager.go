package processor

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"log"
	"time"

	"github.com/arachnys/athenapdf/pkg/proto"
)

var (
	defaultMaxWorkers = 8
	defaultMaxQueue   = 24
	defaultAddTimeout = time.Second * 30
)

type ReadWorkQueue <-chan Work
type WriteWorkQueue chan<- Work

type Manager struct {
	wq         WriteWorkQueue
	addTimeout time.Duration
}

type Worker struct {
	id int
}

type Work struct {
	ctx       context.Context
	processor Processor
	process   *proto.Process

	OutputReader   chan io.Reader
	OutputUploaded chan bool
	Err            chan error
}

func NewManager(maxWorkers, maxQueue int, addTimeout time.Duration) Manager {
	if maxWorkers == 0 {
		maxWorkers = defaultMaxWorkers
	}
	if maxQueue == 0 {
		maxQueue = defaultMaxQueue
	}
	if addTimeout == 0 {
		addTimeout = defaultAddTimeout
	}

	wq := make(chan Work, maxQueue)

	for i := 0; i < maxWorkers; i++ {
		go func(wq ReadWorkQueue, w Worker) {
			for work := range wq {
				log.Printf(
					"worker=%d queue=%d fetcher=%s converter=%s uploader=%s",
					w.id,
					len(wq),
					work.process.GetFetcher(),
					work.process.GetConverter(),
					work.process.GetUploader(),
				)
				work.Start()
			}
		}(ReadWorkQueue(wq), Worker{i})
	}

	return Manager{
		WriteWorkQueue(wq),
		addTimeout,
	}
}

func (pm Manager) Add(ctx context.Context, processor Processor, process *proto.Process) (Work, error) {
	w := Work{
		ctx:       ctx,
		processor: processor,
		process:   process,

		OutputReader:   make(chan io.Reader, 1),
		OutputUploaded: make(chan bool, 1),
		Err:            make(chan error, 1),
	}

	select {
	case <-time.After(pm.addTimeout):
		return Work{}, ProcessError{errors.New("timed out waiting to add process to work queue")}
	case pm.wq <- w:
	}

	return w, nil
}

func (w Work) Start() {
	processCh := make(chan struct {
		r        io.Reader
		uploaded bool
		err      error
	}, 1)

	go func() {
		r, uploaded, err := w.processor.Process(w.ctx, w.process)
		processCh <- struct {
			r        io.Reader
			uploaded bool
			err      error
		}{r, uploaded, err}
	}()

	select {
	case <-w.ctx.Done():
		// Exit, and free up a worker quickly if the context has
		// been cancelled.
		w.Err <- w.ctx.Err()
	case res := <-processCh:
		if res.err != nil {
			w.Err <- res.err
			return
		}

		if res.uploaded {
			w.OutputUploaded <- true
			return
		}

		w.OutputReader <- res.r
	case <-time.After(time.Minute * 5):
		// Prevent leaking goroutine.
		// It is unlikely to happen if the context is correctly used,
		// but timeout just in case there is a rogue subprocess
		// implementation.
		w.Err <- ProcessError{errors.New("timed out waiting for process to complete")}
	}
}

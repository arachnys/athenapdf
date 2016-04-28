package converter

import (
	"errors"
	"runtime"
	"testing"
	"time"
)

func TestInitWorkers(t *testing.T) {
	i := runtime.NumGoroutine()
	wq := InitWorkers(10, 10, 10)
	if wq == nil {
		t.Fatalf("expected work queue to be initialised, not nil")
	}
	defer close(wq)
	if got, want := runtime.NumGoroutine(), i+10; got != want {
		t.Errorf("number of running goroutines is %+v, want %+v", got, want)
	}
}

type TestConversion struct {
	Conversion
}

func (c TestConversion) Convert(s ConversionSource, done <-chan struct{}) ([]byte, error) {
	return []byte("test work"), nil
}

func TestNewWork(t *testing.T) {
	wq := InitWorkers(10, 10, 10)
	c := TestConversion{}
	s := ConversionSource{}
	w := NewWork(wq, c, s)
	if w.out == nil {
		t.Fatalf("expected work output channel to be initialised, not nil")
	}
	if w.err == nil {
		t.Fatalf("expected work error channel to be initialised, not nil")
	}
	if w.uploaded == nil {
		t.Fatalf("expected work uploaded channel to be initialised, not nil")
	}
	if w.done == nil {
		t.Fatalf("expected work done channel to be initialised, not nil")
	}
}

type TestConversionUpload struct {
	Conversion
}

func (c TestConversionUpload) Upload(b []byte) (bool, error) {
	return true, nil
}

func TestNewWork_upload(t *testing.T) {
	wq := InitWorkers(10, 10, 10)
	defer close(wq)
	c := TestConversionUpload{}
	s := ConversionSource{}
	w := NewWork(wq, c, s)
	select {
	case <-w.Uploaded():
	case <-time.After(time.Second):
		t.Errorf("expected work uploaded channel to be closed before timeout")
	}
}

type TestConversionError struct {
	Conversion
}

var (
	ErrTestConversionError = errors.New("test conversion error")
)

func (c TestConversionError) Convert(s ConversionSource, done <-chan struct{}) ([]byte, error) {
	return []byte{}, ErrTestConversionError
}

func TestNewWork_error(t *testing.T) {
	wq := InitWorkers(10, 10, 10)
	defer close(wq)
	c := TestConversionError{}
	s := ConversionSource{}
	w := NewWork(wq, c, s)
	select {
	case <-w.Error():
	case <-time.After(time.Second):
		t.Errorf("expected to receive error before timeout")
	}
}

type TestConversionTimeout struct {
	Conversion
}

func (c TestConversionTimeout) Convert(s ConversionSource, done <-chan struct{}) ([]byte, error) {
	time.Sleep(time.Second * 2)
	return []byte("test work timeout"), nil
}

func TestNewWork_timeout(t *testing.T) {
	wq := InitWorkers(10, 10, 1)
	defer close(wq)
	c := TestConversionTimeout{}
	s := ConversionSource{}
	w := NewWork(wq, c, s)
	go func(w Work) {
		time.Sleep(time.Second * 2)
		close(w.err)
	}(w)
	if got, want := <-w.Error(), ErrConversionTimeout; got != want {
		t.Errorf("expected a conversion timeout error, got %+v", got)
	}
}

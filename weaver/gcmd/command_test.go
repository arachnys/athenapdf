package gcmd

import (
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	testString := "test execute"
	mockTerminate := make(chan struct{}, 1)
	got, err := Execute([]string{"echo", testString}, mockTerminate)
	if err != nil {
		t.Fatalf("execute returned an unexpected error: %+v", err)
	}
	if want := []byte(testString + "\n"); !reflect.DeepEqual(got, want) {
		t.Errorf("expected output of executed command to be %+v, got %+v", want, got)
	}
}

func TestExecute_err(t *testing.T) {
	testString := "test execute"
	mockTerminate := make(chan struct{}, 1)
	got, err := Execute([]string{"echo-broken", testString}, mockTerminate)
	if err == nil {
		t.Fatalf("expected error to be returned")
	}
	if got != nil {
		t.Errorf("expected output of executed command to be nil, got %+v", got)
	}
}

func TestExecute_done(t *testing.T) {
	testString := "test execute"
	mockTerminate := make(chan struct{}, 1)
	close(mockTerminate)
	got, err := Execute([]string{"echo", testString}, mockTerminate)
	if err == nil {
		t.Fatalf("expected error to be returned")
	}
	if err != ErrCmdTerminated {
		t.Errorf("expected a command terminated error")
	}
	if got != nil {
		t.Errorf("expected output of executed command to be nil, got %+v", got)
	}
}

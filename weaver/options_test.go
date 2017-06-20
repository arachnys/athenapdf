package main

import (
	"testing"

	"gopkg.in/guregu/null.v3"
	"strings"
)

var nilResult error = nil

func TestRuntimeOptionsEmpty(t *testing.T) {

	ro := &RuntimeOptions{}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be %d, got %d", want, got)
	}
}

func TestRuntimeOptionsPageSize(t *testing.T) {

	ro := &RuntimeOptions{Pagesize: null.StringFrom("asdf")}
	if got, want := ro.Validate(), RuntimeOptionsPageSizeError; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Pagesize: null.StringFrom("a3")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--pagesize A3"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Pagesize: null.StringFrom("a4")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--pagesize A4"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Pagesize: null.StringFrom("a5")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--pagesize A5"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Pagesize: null.StringFrom("letter")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--pagesize Letter"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Pagesize: null.StringFrom("tabloid")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--pagesize Tabloid"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
}

func TestRuntimeOptionsMargins(t *testing.T) {

	ro := &RuntimeOptions{Margins: null.StringFrom("asdf")}
	if got, want := ro.Validate(), RuntimeOptionsMarginsError; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Margins: null.StringFrom("standard")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--margins standard"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Margins: null.StringFrom("none")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--margins none"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}

	ro = &RuntimeOptions{Margins: null.StringFrom("minimal")}
	if got, want := ro.Validate(), nilResult; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
	if got, want := strings.Join(ro.BuildCommand(), " "), "--margins minimal"; got != want {
		t.Fatalf("expected validate to be (%v), got (%v)", want, got)
	}
}

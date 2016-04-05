package util

import (
	"github.com/arachnys/athenapdf/weaver/testutil"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestHandleOctetStream(t *testing.T) {
	testString := "test application/octet-stream bytes"
	ts := testutil.MockHTTPServer("application/octet-stream", testString)
	defer ts.Close()
	tmp, err := HandleOctetStream(ts.URL)
	if err != nil {
		t.Fatalf("handleoctetstream returned an unexpected error: %+v", err)
	}
	if tmp == "" {
		t.Errorf("expected tmp file path to be returned")
	}
	defer os.Remove(tmp)
	got, err := ioutil.ReadFile(tmp)
	if err != nil {
		t.Fatalf("unable to read tmp file: %+v", err)
	}
	if want := []byte(testString); !reflect.DeepEqual(got, want) {
		t.Errorf("expected created tmp file bytes to be %+v, got %+v", want, got)
	}
}

func TestHandleOctetStream_html(t *testing.T) {
	ts := testutil.MockHTTPServer("", "should not be saved")
	defer ts.Close()
	got, err := HandleOctetStream(ts.URL)
	if err != nil {
		t.Fatalf("handleoctetstream returned an unexpected error: %+v", err)
	}
	if got != "" {
		t.Errorf("expected no tmp file path, got %+v", got)
	}
}

func TestHandleOctetStream_badURL(t *testing.T) {
	got, err := HandleOctetStream("bad URL")
	if err == nil {
		t.Fatalf("expected error to be returned")
	}
	if got != "" {
		t.Errorf("expected no tmp file path, got %+v", got)
	}
}

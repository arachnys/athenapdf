package converter

import (
	"bytes"
	"github.com/arachnys/athenapdf/weaver/testutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func expectLocalConversion(t *testing.T, s *ConversionSource, uri string, mime string, data []byte) {
	// Test if URI is overwritten
	if uri == s.URI {
		t.Fatalf("expected conversion source URI to be overwritten with a local file path")
	}
	// Test if file path is absolute
	if !filepath.IsAbs(s.URI) {
		t.Fatalf("expected conversion source URI to be an absolute temporary file path, got %s", s.URI)
	}
	// GC test file
	defer os.Remove(s.URI)
	// Test content type
	if got, want := s.Mime, mime; got != want {
		t.Errorf("expected content type of conversion source to be %s, got %s", want, got)
	}
	// Test if IsLocal flag is set
	if !s.IsLocal {
		t.Errorf("expected IsLocal to be set to true for raw source conversions")
	}
	// Test file contents
	got, err := ioutil.ReadFile(s.URI)
	if err != nil {
		t.Fatalf("unable to read temporary file: %+v", err)
	}
	if want := data; !reflect.DeepEqual(got, want) {
		t.Errorf("expected created temporary file bytes to be %+v, got %+v", want, got)
	}
}

func TestReaderContentType(t *testing.T) {
	mockReader := strings.NewReader("<!DOCTYPE HTML>")
	got, err := readerContentType(mockReader)
	if err != nil {
		t.Fatalf("readercontenttype returned an unexpected error: %+v", err)
	}
	if want := "text/html; charset=utf-8"; got != want {
		t.Errorf("expected content type of byte stream to be %s, got %s", want, got)
	}
}

func TestReaderContentType_eof(t *testing.T) {
	mockReader := new(bytes.Buffer)
	got, err := readerContentType(mockReader)
	if err != nil {
		t.Fatalf("readercontenttype returned an unexpected error: %+v", err)
	}
	if want := "text/plain; charset=utf-8"; got != want {
		t.Errorf("expected content type of byte stream to be %s, got %s", want, got)
	}
}

func TestReaderTmpFile(t *testing.T) {
	mockData := "<!DOCTYPE HTML>"
	mockReader := strings.NewReader(mockData)
	fp, ft, err := readerTmpFile(mockReader)
	if err != nil {
		t.Fatalf("readertmpfile returned an unexpected error: %+v", err)
	}
	if !filepath.IsAbs(fp) {
		t.Errorf("expected temporary file path to be absolute, got %s", fp)
	}
	defer os.Remove(fp)
	got, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("unable to read temporary file: %+v", err)
	}
	if want := []byte(mockData); !reflect.DeepEqual(got, want) {
		t.Errorf("expected created temporary file bytes to be %+v, got %+v", want, got)
	}
	if got, want := ft, "text/html; charset=utf-8"; got != want {
		t.Errorf("expected content type of byte stream to be %s, got %s", want, got)
	}
}

func TestRawSource(t *testing.T) {
	s := new(ConversionSource)
	mockURI := "http://this-should-be-overwritten"
	s.URI = mockURI
	mockData := "<!DOCTYPE HTML>"
	mockReader := strings.NewReader(mockData)
	err := rawSource(s, mockReader)
	if err != nil {
		t.Fatalf("rawsource returned an unexpected error: %+v", err)
	}
	expectLocalConversion(t, s, mockURI, "text/html; charset=utf-8", []byte(mockData))
}

func TestUriSource(t *testing.T) {
	s := new(ConversionSource)
	ts := testutil.MockHTTPServer("", "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	defer ts.Close()
	err := uriSource(s, ts.URL)
	if err != nil {
		t.Fatalf("urisource returned an unexpected error: %+v", err)
	}
	if got, want := s.URI, ts.URL; got != want {
		t.Errorf("expected URI of conversion source to be %s, got %s", want, got)
	}
	if got, want := s.Mime, "text/xml; charset=utf-8"; got != want {
		t.Errorf("expected content type of conversion source to be %s, got %s", want, got)
	}
	if got := s.OriginalURI; got != "" {
		t.Errorf("expected original URI of conversion source to be empty, got %s", got)
	}
}

func TestUriSource_octet(t *testing.T) {
	s := new(ConversionSource)
	mockData := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	ts := testutil.MockHTTPServer("application/octet-stream", mockData)
	defer ts.Close()
	err := uriSource(s, ts.URL)
	if err != nil {
		t.Fatalf("urisource returned an unexpected error: %+v", err)
	}
	expectLocalConversion(t, s, ts.URL, "text/xml; charset=utf-8", []byte(mockData))
	if got, want := s.OriginalURI, ts.URL; got != want {
		t.Errorf("expected original URI of conversion source to be %s, got %s", want, got)
	}
}

func TestNewConversionSource(t *testing.T) {
	mockURI := "http://this-should-not-be-used"
	mockData := "<!DOCTYPE HTML>"
	mockReader := strings.NewReader(mockData)
	s, err := NewConversionSource(mockURI, mockReader)
	if err != nil {
		t.Fatalf("newconversionsource returned an unexpected error: %+v", err)
	}
	expectLocalConversion(t, s, mockURI, "text/html; charset=utf-8", []byte(mockData))
}

func TestNewConversionSource_remote(t *testing.T) {
	ts := testutil.MockHTTPServer("", "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	defer ts.Close()
	s, err := NewConversionSource(ts.URL, nil)
	if err != nil {
		t.Fatalf("newconversionsource returned an unexpected error: %+v", err)
	}
	if got, want := s.URI, ts.URL; got != want {
		t.Errorf("expected URI of conversion source to be %s, got %s", want, got)
	}
	if got, want := s.Mime, "text/xml; charset=utf-8"; got != want {
		t.Errorf("expected content type of conversion source to be %s, got %s", want, got)
	}
	if got := s.OriginalURI; got != "" {
		t.Errorf("expected original URI of conversion source to be empty, got %s", got)
	}
}

func TestGetActualURI(t *testing.T) {
	s := new(ConversionSource)
	mockURI := "http://this-should-be-returned"
	s.OriginalURI = mockURI
	s.URI = "http://this-should-not-be-returned"
	if got, want := s.GetActualURI(), mockURI; got != want {
		t.Errorf("expected the original conversion target to be %s, got %s", want, got)
	}
}

func TestGetActualURI_noOriginalURI(t *testing.T) {
	s := new(ConversionSource)
	mockURI := "http://this-should-be-returned"
	s.URI = mockURI
	if got, want := s.GetActualURI(), mockURI; got != want {
		t.Errorf("expected the original conversion target to be %s, got %s", want, got)
	}
}

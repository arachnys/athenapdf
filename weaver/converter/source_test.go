package converter

import (
	"bytes"
	"github.com/arachnys/athenapdf/weaver/testutil"
	"io/ioutil"
	"net/url"
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
	// Test if IsLocal flag is not set
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

func expectRemoteConversion(t *testing.T, s *ConversionSource, uri string, mime string) {
	// Test if URI matches
	if got, want := s.URI, uri; got != want {
		t.Errorf("expected URI of conversion source to be %s, got %s", want, got)
	}
	// Test content type
	if got, want := s.Mime, mime; got != want {
		t.Errorf("expected content type of conversion source to be %s, got %s", want, got)
	}
	// Test if original URI is set
	if got := s.OriginalURI; got != "" {
		t.Errorf("expected original URI of conversion source to be empty, got %s", got)
	}
	// Test if IsLocal flag is set
	if s.IsLocal {
		t.Errorf("expected IsLocal to be set to false for remote conversions")
	}
}

func setMockURI(t *testing.T, s *ConversionSource) {
	mockFile, err := ioutil.TempFile(os.TempDir(), "tmp")
	if err != nil {
		t.Fatalf("unable to create temporary mock file: %+v", err)
	}

	p, err := filepath.Abs(mockFile.Name())
	if err != nil {
		t.Fatalf("unable to create absolute representation of mock file path: %+v", err)
	}

	s.URI = p
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
	defer os.Remove(fp)
	if !filepath.IsAbs(fp) {
		t.Errorf("expected temporary file path to be absolute, got %s", fp)
	}
	if got, want := ft, "text/html; charset=utf-8"; got != want {
		t.Errorf("expected content type of byte stream to be %s, got %s", want, got)
	}
	got, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("unable to read temporary file: %+v", err)
	}
	if want := []byte(mockData); !reflect.DeepEqual(got, want) {
		t.Errorf("expected created temporary file bytes to be %+v, got %+v", want, got)
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
	ts := testutil.MockHTTPServer("", "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", false)
	defer ts.Close()
	err := uriSource(s, ts.URL)
	if err != nil {
		t.Fatalf("urisource returned an unexpected error: %+v", err)
	}
	expectRemoteConversion(t, s, ts.URL, "text/xml; charset=utf-8")
}

func TestUriSource_basicAuth(t *testing.T) {
	s := new(ConversionSource)
	ts := testutil.MockHTTPServer("", "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", true)
	defer ts.Close()

	// Test unauthenticated
	err := uriSource(s, ts.URL)
	if err != nil {
		t.Fatalf("urisource (unauthenticated) returned an unexpected error: %+v", err)
	}
	// Unauthenticated content type (wrong)
	if got, want := s.Mime, "text/plain; charset=utf-8"; got != want {
		t.Errorf("expected content type of conversion source to be %s, got %s", want, got)
	}

	// Test authenticated
	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("failed to parse mock server url: %+v", err)
	}

	u.User = url.UserPassword("test", "test")
	err = uriSource(s, u.String())
	if err != nil {
		t.Fatalf("urisource (authenticated) returned an unexpected error: %+v", err)
	}
	// Authenticated content type (correct)
	if got, want := s.Mime, "text/xml; charset=utf-8"; got != want {
		t.Errorf("expected content type of conversion source to be %s, got %s", want, got)
	}
}

func TestUriSource_octet(t *testing.T) {
	s := new(ConversionSource)
	mockData := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
	ts := testutil.MockHTTPServer("application/octet-stream", mockData, false)
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

func TestSetCustomExtension(t *testing.T) {
	s := new(ConversionSource)
	setMockURI(t, s)
	defer os.Remove(s.URI)

	s.IsLocal = true
	original := s.URI

	err := setCustomExtension(s, "mhtml")
	if err != nil {
		t.Fatalf("setcustomextension returned an unexpected error: %+v", err)
	}

	if got, want := s.URI, original+".mhtml"; got != want {
		t.Errorf("expected conversion source URI to be %s, got %s", want, got)
	}
}

func TestSetCustomExtension_error(t *testing.T) {
	s := new(ConversionSource)
	s.URI = os.TempDir() + "/doesnotexist"
	s.IsLocal = true

	err := setCustomExtension(s, "mhtml")
	if err == nil {
		t.Fatalf("expected error to be returned")
	}
}

func TestSetCustomExtension_notLocal(t *testing.T) {
	s := new(ConversionSource)
	setMockURI(t, s)
	defer os.Remove(s.URI)

	original := s.URI

	err := setCustomExtension(s, "mhtml")
	if err != nil {
		t.Fatalf("setcustomextension returned an unexpected error: %+v", err)
	}

	if got, want := s.URI, original; got != want {
		t.Errorf("expected conversion source URI to be %s, got %s", want, got)
	}
}

func TestSetCustomExtension_noExt(t *testing.T) {
	s := new(ConversionSource)
	setMockURI(t, s)
	defer os.Remove(s.URI)

	s.IsLocal = true
	original := s.URI

	err := setCustomExtension(s, "")
	if err != nil {
		t.Fatalf("setcustomextension returned an unexpected error: %+v", err)
	}

	if got, want := s.URI, original; got != want {
		t.Errorf("expected conversion source URI to be %s, got %s", want, got)
	}
}

func TestNewConversionSource(t *testing.T) {
	mockURI := "http://this-should-not-be-used"
	mockData := "<!DOCTYPE HTML>"
	mockReader := strings.NewReader(mockData)
	s, err := NewConversionSource(mockURI, mockReader, "")
	if err != nil {
		t.Fatalf("newconversionsource returned an unexpected error: %+v", err)
	}
	expectLocalConversion(t, s, mockURI, "text/html; charset=utf-8", []byte(mockData))
}

func TestNewConversionSource_remote(t *testing.T) {
	ts := testutil.MockHTTPServer("", "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", false)
	defer ts.Close()
	s, err := NewConversionSource(ts.URL, nil, "")
	if err != nil {
		t.Fatalf("newconversionsource returned an unexpected error: %+v", err)
	}
	expectRemoteConversion(t, s, ts.URL, "text/xml; charset=utf-8")
}

func TestNewConversionSource_invalidURL(t *testing.T) {
	s, err := NewConversionSource("http://invalid-url", nil, "")
	if err == nil {
		t.Fatalf("expected error to be returned")
	}
	if s != nil {
		t.Fatalf("expected result of newconversionsource to be nil, got %+v", s)
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

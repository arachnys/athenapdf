package converter

import (
	"golang.org/x/net/publicsuffix"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
)

// ConversionSource contains the target resource path, and its MIME type.
// It may contain additional metadata about the conversion source.
type ConversionSource struct {
	// URI represents the actual path to a remote (web) or local resource
	// to be converted.
	URI string
	// OriginalURI may contain the original path to a remote resource to be
	// converted if it is downloaded for local conversion.
	// This normally happens when the resource is an `octet-stream`.
	// It is only metadata, and it should NOT be used as the conversion source.
	OriginalURI string
	// Mime represents the content type of the source. It will be set to
	// `application/octet-stream` if it is unable to determine the type.
	// For more information see:
	// https://mimesniff.spec.whatwg.org/#matching-a-mime-type-pattern
	Mime string
	// IsLocal should be true if the URI relies on a local conversion strategy,
	// and false if the target is a remote source (that does not require
	// pre-processing).
	IsLocal bool
}

// readerContentType attempts to determine the content type using bytes from a
// reader. It returns the content type if successful or an empty string,
// and an error if unsuccessful.
func readerContentType(r io.Reader) (string, error) {
	// Only the first 512 bytes of data is needed to determine the content type
	b := make([]byte, 512)
	n, err := r.Read(b)
	// Ignore cases where the file size is smaller than 512 bytes
	if err != nil && err != io.EOF {
		return "", err
	}
	return http.DetectContentType(b[:n]), nil
}

// readerTmpFile creates a temporary file using bytes from a reader.
// It returns the full temporary file path, and its mime type if successful.
func readerTmpFile(r io.Reader) (string, string, error) {
	// Create a temporary file
	f, err := ioutil.TempFile(os.TempDir(), "tmp")
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	// Pipe bytes from a reader to a file writer
	if _, err = io.Copy(f, r); err != nil {
		return "", "", err
	}

	// Determine full temporary file path
	p, err := filepath.Abs(f.Name())
	if err != nil {
		return "", "", err
	}

	// Reset read / write offset to 0
	_, err = f.Seek(0, 0)
	if err != nil {
		return "", "", err
	}

	// Determine file content type from file reader
	t, err := readerContentType(f)
	if err != nil {
		return "", "", err
	}

	return p, t, nil
}

// rawSource is a local conversion strategy handler. It accepts a reader,
// and performs pre-processing to save the stream of bytes to a local file
// for conversion.
func rawSource(s *ConversionSource, body io.Reader) error {
	// Save content locally in a temporary file
	p, t, err := readerTmpFile(body)
	if err != nil {
		return err
	}
	s.URI = p
	s.Mime = t
	s.IsLocal = true
	return nil
}

// uriSource is a remote conversion strategy handler. It will attempt to fetch
// the remote URI to determine: if it is accessible; its mime type; and
// if it needs pre-processing (e.g. `octet-stream`).
func uriSource(s *ConversionSource, uri string) error {
	// Fetch URL with support for cookies (to handle session-based redirects)
	opts := cookiejar.Options{PublicSuffixList: publicsuffix.List}
	jar, err := cookiejar.New(&opts)
	if err != nil {
		return err
	}
	client := http.Client{Jar: jar}
	res, err := client.Get(uri)
	if err != nil {
		return err
	}
	if res != nil {
		defer res.Body.Close()
	}

	// Save content locally (temporarily) if the HTTP header indicates that it
	// is a binary stream.
	// TODO: file restrictions / limits (e.g. size)
	if res.Header.Get("Content-Type") == "application/octet-stream" {
		// Set the OriginalURI as we are running a local conversion strategy
		s.OriginalURI = uri
		// Pipe HTTP response body to a temporary file via io.Reader
		if rawSource(s, res.Body) != nil {
			return err
		}
	} else {
		// Do not set the OriginalURI as it is NOT a local conversion
		s.URI = uri
		// Read the first 512 bytes of the page contents into a temporary buffer
		// so that we can determine the content type
		t, err := readerContentType(res.Body)
		if err != nil {
			return err
		}
		s.Mime = t
	}

	return nil
}

func setCustomExtension(s *ConversionSource, ext string) error {
	if s.IsLocal && len(ext) > 0 {
		newPath := s.URI + "." + ext
		if err := os.Rename(s.URI, newPath); err != nil {
			return err
		}
		s.URI = newPath
	}
	return nil
}

// NewConversionSource creates, and returns a new ConversionSource.
// It accepts either a URI to a remote resource or a reader containing a stream
// of bytes. If both parameters are specified, the reader takes precedence.
// The ConversionSource is prepared using one of two strategies: a local
// conversion (see rawSource) or a remote conversion (see uriSource).
func NewConversionSource(uri string, body io.Reader, ext string) (*ConversionSource, error) {
	s := new(ConversionSource)

	var err error
	if body != nil {
		err = rawSource(s, body)
	} else {
		err = uriSource(s, uri)
	}

	if err != nil {
		return nil, err
	}

	if err := setCustomExtension(s, ext); err != nil {
		return nil, err
	}

	return s, nil
}

// GetActualURI returns the original conversion target. The URI field may
// not contain the original conversion target as it may be overwritten when
// using a local conversion strategy.
func (s ConversionSource) GetActualURI() string {
	uri := s.URI
	if s.OriginalURI != "" {
		uri = s.OriginalURI
	}
	return uri
}

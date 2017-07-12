package libreoffice

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/arachnys/athenapdf/pkg/proto"
)

func TestConvert(t *testing.T) {
	testInput, err := filepath.Abs("../testdata/test.docx")
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	u, err := url.Parse(testInput)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
	u.Scheme = "file"

	converter := &LibreOffice{}
	ctx := context.TODO()
	conversion := &proto.Conversion{
		Uri: u.String(),
	}

	gotReader, err := converter.Convert(ctx, conversion, nil)
	if err != nil {
		t.Fatalf("failed to convert, unexpected error: %+v", err)
	}

	gotData, err := ioutil.ReadAll(gotReader)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	if !bytes.Contains(gotData, []byte("%PDF-1.4")) {
		t.Errorf("output is not a valid pdf: %+v", gotData)
	}
}

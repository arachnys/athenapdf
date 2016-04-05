package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestNewTmpFile(t *testing.T) {
	testString := "test string"
	mockReader := strings.NewReader(testString)
	mockReadCloser := ioutil.NopCloser(mockReader)
	tmp, err := NewTmpFile(mockReadCloser)
	if err != nil {
		t.Fatalf("newtmpfile returned an unexpected error: %+v", err)
	}
	defer os.Remove(tmp)
	if !filepath.IsAbs(tmp) {
		t.Errorf("expected tmp file path to be absolute, got %+v", tmp)
	}
	got, err := ioutil.ReadFile(tmp)
	if err != nil {
		t.Fatalf("unable to read tmp file: %+v", err)
	}
	if want := []byte(testString); !reflect.DeepEqual(got, want) {
		t.Errorf("expected created tmp file bytes to be %+v, got %+v", want, got)
	}
}

package converter

import (
	"reflect"
	"testing"
)

func TestConversion_Convert(t *testing.T) {
	mockConversion := Conversion{}
	mockSource := ConversionSource{}
	mockDone := make(chan struct{}, 1)
	got, err := mockConversion.Convert(mockSource, mockDone)
	if err != nil {
		t.Fatalf("convert returned an unexpected error: %+v", err)
	}
	if want := []byte{}; !reflect.DeepEqual(got, want) {
		t.Errorf("expected output of conversion to be %+v, got %+v", want, got)
	}
}

func TestConversion_Upload(t *testing.T) {
	mockConversion := Conversion{}
	got, err := mockConversion.Upload(nil)
	if err != nil {
		t.Fatalf("upload returned an unexpected error: %+v", err)
	}
	if want := false; got != want {
		t.Errorf("expected status of upload to be %+v, got %+v", want, got)
	}
}

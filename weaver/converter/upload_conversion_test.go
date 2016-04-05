package converter

import (
	"testing"
)

func expectUploadToHalt(t *testing.T, mockConversion UploadConversion) {
	got, err := mockConversion.Upload([]byte{})
	if err != nil {
		t.Fatalf("upload returned an unexpected error: %+v", err)
	}
	if want := false; got != want {
		t.Errorf("expected status of upload to be %+v, got %+v", want, got)
	}
}

func TestUploadConversion_Upload_noS3Bucket(t *testing.T) {
	mockConversion := UploadConversion{}
	mockConversion.AWSS3.S3Key = "s3-key-123456"
	expectUploadToHalt(t, mockConversion)
}

func TestUploadConversion_Upload_noS3Key(t *testing.T) {
	mockConversion := UploadConversion{}
	mockConversion.AWSS3.S3Bucket = "s3-bucket-123456"
	expectUploadToHalt(t, mockConversion)
}

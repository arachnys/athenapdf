package cloudconvert

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewProcess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("unable to read request body: %+v", err)
		}
		want := "apikey=test+cloudconvert+key&inputformat=html&outputformat=pdf"
		if got := string(b); got != want {
			t.Errorf("expected cloudconvert process request body to be %s, got %+v", want, got)
		}
		mockReturn := `{"url":"\/\/test.cloudconvert.com\/process\/123123","id":"123123","host":"test.cloudconvert.com","expires":"2016-03-08 16:06:25","maxtime":18000,"minutes":4258}`
		w.Write([]byte(mockReturn))
		w.Header().Set("Content-Type", "application/json")
	}))
	defer ts.Close()

	c := Client{
		ts.URL,
		"test cloudconvert key",
		0,
	}
	process, err := c.NewProcess("html", "pdf")
	if err != nil {
		t.Fatalf("newprocess returned an unexpected error: %+v", err)
	}
	if got, want := process.URL, "https://test.cloudconvert.com/process/123123"; got != want {
		t.Errorf("expected cloudconvert process URL to be %s, got %+v", want, got)
	}
}

func TestStartConversion(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var got Conversion
		if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
			t.Fatalf("unable to json decode request body: %+v", err)
		}
		s3 := S3{
			AccessKey:    "aws id",
			AccessSecret: "aws secret",
			Bucket:       "s3 bucket",
			Path:         "s3 key",
			ACL:          "public-read",
		}
		output := Output{s3}
		want := Conversion{
			Input:        "download",
			File:         "http://www.test-url.com/",
			Filename:     "s3 key.html",
			OutputFormat: "pdf",
			Wait:         true,
			Output:       &output,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected cloudconvert conversion request body to be %+v, got %+v", want, got)
		}
		w.Header().Set("Content-Type", "application/json")
	}))
	defer ts.Close()

	p := Process{
		URL: ts.URL,
	}
	output := Output{
		S3{
			"aws id",
			"aws secret",
			"s3 bucket",
			"s3 key",
			"public-read",
		},
	}
	c := Conversion{
		Input:        "download",
		File:         "http://www.test-url.com/",
		Filename:     "s3 key.html",
		OutputFormat: "pdf",
		Wait:         true,
		Output:       &output,
	}
	got, err := p.StartConversion(c)
	if err != nil {
		t.Fatalf("startconversion returned an unexpected error: %+v", err)
	}
	if got != nil {
		t.Errorf("expected output of startconversion to be nil, got %+v", got)
	}
}

func expectUploadStatus(t *testing.T, c CloudConvert, want bool) {
	got, err := c.Upload(nil)
	if err != nil {
		t.Fatalf("upload returned an unexpected error: %+v", err)
	}
	if got != want {
		t.Errorf("expected status of upload to be %+v, got %+v", want, got)
	}
}

func TestUpload(t *testing.T) {
	c := CloudConvert{}
	expectUploadStatus(t, c, false)
}

func TestUpload_noS3Bucket(t *testing.T) {
	c := CloudConvert{}
	c.AWSS3.S3Key = "s3-key-123456"
	expectUploadStatus(t, c, false)
}

func TestUpload_noS3Key(t *testing.T) {
	c := CloudConvert{}
	c.AWSS3.S3Bucket = "s3-bucket-123456"
	expectUploadStatus(t, c, false)
}

func TestUpload_s3(t *testing.T) {
	c := CloudConvert{}
	c.AWSS3.S3Bucket = "s3-bucket-123456"
	c.AWSS3.S3Key = "s3-key-123456"
	expectUploadStatus(t, c, true)
}

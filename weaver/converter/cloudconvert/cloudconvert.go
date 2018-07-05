package cloudconvert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/arachnys/athenapdf/weaver/converter"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CloudConvert struct {
	converter.UploadConversion
	Client
}

type Client struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

type Process struct {
	ID      string `json:"id,omitempty"`
	URL     string `json:"url"`
	Expires string `json:"expires,omitempty"`
	MaxTime int    `json:"maxtime,omitempty"`
	Minutes int    `json:"minutes,omitempty"`
}

type S3 struct {
	AccessKey    string `json:"accesskeyid"`
	AccessSecret string `json:"secretaccesskey"`
	Bucket       string `json:"bucket"`
	Path         string `json:"path"`
	ACL          string `json:"acl"`
}

type Output struct {
	S3 `json:"s3"`
}

type Conversion struct {
	Input        string `json:"input"`
	File         string `json:"file"`
	Filename     string `json:"filename"`
	OutputFormat string `json:"outputformat"`
	Wait         bool   `json:"wait"`
	Download     string `json:"download,omitempty"`
	Timeout      string `json:"timeout,omitempty"`
	*Output      `json:"output,omitempty"`
}

func (c Client) QuickConversion(path string, awsS3 converter.AWSS3, inputFormat string, outputFormat string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b := new(bytes.Buffer)
	bw := multipart.NewWriter(b)

	// Default timeout: 5 minutes
	if c.Timeout == 0 {
		c.Timeout = time.Minute * 5
	}

	// Use a map so we can easily extend the parameters (options)
	params := map[string]string{
		"apikey":       c.APIKey,
		"input":        "upload",
		"download":     "inline",
		"filename":     "tmp.html",
		"inputformat":  inputFormat,
		"outputformat": outputFormat,
		"timeout":      fmt.Sprintf("%.0f", c.Timeout.Seconds()),
	}

	part, err := bw.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, f)

	for k, v := range params {
		err = bw.WriteField(k, v)
		if err != nil {
			return nil, err
		}
	}

	err = bw.Close()
	if err != nil {
		return nil, err
	}

	res, err := http.Post(c.BaseURL+"/convert", bw.FormDataContentType(), b)
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		var data map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("[CloudConvert] did not receive HTTP 200, response: %+v\n", data)
	}

	o, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (c Client) NewProcess(inputFormat, ouputFormat string) (Process, error) {
	process := Process{}
	res, err := http.PostForm(
		c.BaseURL+"/process",
		url.Values{
			"apikey":       {c.APIKey},
			"inputformat":  {inputFormat},
			"outputformat": {ouputFormat},
		},
	)
	if err != nil {
		return process, err
	}
	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		var data map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
			return process, err
		}
		return process, fmt.Errorf("[CloudConvert] did not receive HTTP 200, response: %+v\n", data)
	}

	err = json.NewDecoder(res.Body).Decode(&process)
	if err == nil && strings.HasPrefix(process.URL, "//") {
		process.URL = "https:" + process.URL
	}

	return process, err
}

func (p Process) StartConversion(c Conversion) ([]byte, error) {
	b, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}
	res, err := http.Post(p.URL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	if res != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != 200 {
		var data map[string]interface{}
		if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("[CloudConvert] did not receive HTTP 200, response: %+v\n", data)
	}

	if c.Download == "inline" {
		o, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return o, nil
	}

	return nil, nil
}

func (c CloudConvert) Convert(s converter.ConversionSource, done <-chan struct{}) ([]byte, error) {
	log.Printf("[CloudConvert] converting to PDF: %s\n", s.GetActualURI())

	var b []byte

	if s.IsLocal {
		b, err := c.Client.QuickConversion(s.URI, c.AWSS3, "html", "pdf")
		if err != nil {
			return nil, err
		}
		return b, nil
	}

	p, err := c.Client.NewProcess("html", "pdf")
	if err != nil {
		return nil, err
	}

	conv := Conversion{
		Input:        "download",
		File:         s.URI,
		Filename:     c.AWSS3.S3Key + ".html",
		OutputFormat: "pdf",
		Wait:         true,
		Timeout:      fmt.Sprintf("%.0f", c.Timeout.Seconds()),
	}

	u := uuid.NewV4()

	if c.AWSS3.S3Bucket == "" || c.AWSS3.S3Key == "" {
		conv.Download = "inline"
		conv.Filename = u.String() + ".html"
	} else {
		conv.Output = &Output{
			S3{
				c.AWSS3.AccessKey,
				c.AWSS3.AccessSecret,
				c.AWSS3.S3Bucket,
				c.AWSS3.S3Key,
				"public-read",
			},
		}
		log.Printf("[CloudConvert] uploading conversion to S3: %s\n", c.AWSS3.S3Key)
	}

	b, err = p.StartConversion(conv)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c CloudConvert) Upload(b []byte) (bool, error) {
	if c.AWSS3.S3Bucket == "" || c.AWSS3.S3Key == "" {
		return false, nil
	}

	if b != nil {
		if _, err := c.UploadConversion.Upload(b); err != nil {
			return false, err
		}
	}

	return true, nil
}

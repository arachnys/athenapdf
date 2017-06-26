package converter

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

type AWSS3 struct {
	Region       string
	AccessKey    string
	AccessSecret string
	S3Bucket     string
	S3Key        string
	S3Acl        string
}

type UploadConversion struct {
	Conversion
	AWSS3
}

func uploadToS3(awsConf AWSS3, b []byte) error {
	log.Printf("[Converter] uploading conversion to S3 bucket '%s' with key '%s'\n", awsConf.S3Bucket, awsConf.S3Key)
	st := time.Now()

	region := "us-east-1"
	if awsConf.Region != "" {
		region = awsConf.Region
	}

	acl := "public-read"
	if awsConf.S3Acl != "" {
		acl = awsConf.S3Acl
	}

	conf := aws.NewConfig().WithRegion(region).WithMaxRetries(3)

	if awsConf.AccessKey != "" && awsConf.AccessSecret != "" {
		creds := credentials.NewStaticCredentials(awsConf.AccessKey, awsConf.AccessSecret, "")

		// Credential 'Value'
		_, err := creds.Get()
		if err != nil {
			return err
		}

		conf = conf.WithCredentials(creds)
	}

	sess := session.New(conf)
	svc := s3.New(sess)

	p := &s3.PutObjectInput{
		Bucket:      aws.String(awsConf.S3Bucket),
		Key:         aws.String(awsConf.S3Key),
		ACL:         aws.String(acl),
		ContentType: aws.String("application/pdf"),
		Body:        bytes.NewReader(b),
	}

	res, err := svc.PutObject(p)
	if err != nil {
		return err
	}

	et := time.Now()
	log.Printf("[Converter] uploaded to S3: %s (%s)\n", awsutil.StringValue(res), et.Sub(st))
	return nil
}

func (c UploadConversion) Upload(b []byte) (bool, error) {
	if c.AWSS3.S3Bucket == "" || c.AWSS3.S3Key == "" {
		return false, nil
	}

	if err := uploadToS3(c.AWSS3, b); err != nil {
		return false, err
	}

	return true, nil
}

package s3

import (
	"context"
	"io"

	"github.com/arachnys/athenapdf/pkg/config"
	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/uploader"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

const (
	uploaderName = "s3"

	defaultAwsRegion = "us-east-1"
	defaultS3Acl     = "public-read"
)

type S3 struct{}

func init() {
	uploader.Register(uploaderName, &S3{})
}

// TODO: add support for cancellations
func (_ *S3) Upload(c context.Context, r io.Reader, opts map[string]*proto.Option) error {
	conf := config.MustGet(uploaderName, opts)

	accessKey := conf("access_key")
	accessSecret := conf("access_secret")
	acl := conf("acl")
	bucket := conf("bucket")
	region := conf("region")

	if acl == "" {
		acl = defaultS3Acl
	}

	if region == "" {
		region = defaultAwsRegion
	}

	awsConf := aws.NewConfig().WithRegion(region).WithMaxRetries(3)

	// Use static credentials if defined
	if accessKey != "" && accessSecret != "" {
		creds := credentials.NewStaticCredentials(accessKey, accessSecret, "")
		if _, err := creds.Get(); err != nil {
			return err
		}

		awsConf = awsConf.WithCredentials(creds)
	}

	svc := awss3.New(session.New(awsConf))
	input := &awss3.PutObjectInput{
		ACL:         aws.String("public-read"),
		Body:        aws.ReadSeekCloser(r),
		Bucket:      aws.String(bucket),
		ContentType: aws.String("application/pdf"),
		Key:         aws.String(acl),
	}

	if _, err := svc.PutObject(input); err != nil {
		return err
	}

	return nil
}

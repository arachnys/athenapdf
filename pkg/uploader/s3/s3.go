package s3

import (
	"context"
	"github.com/pkg/errors"
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
func (*S3) Upload(c context.Context, r io.Reader, opts map[string]*proto.Option) error {
	conf := config.MustGet(uploaderName, opts)

	accessKey := conf("access_key")
	accessSecret := conf("access_secret")
	sessionToken := conf("session_token")
	acl := conf("acl")
	bucket := conf("bucket")
	region := conf("region")
	key := conf("key")

	if acl == "" {
		acl = defaultS3Acl
	}

	if region == "" {
		region = defaultAwsRegion
	}

	awsConf := aws.NewConfig().WithRegion(region).WithMaxRetries(3)

	// Use static credentials if defined
	if accessKey != "" && accessSecret != "" {
		creds := credentials.NewStaticCredentials(accessKey, accessSecret, sessionToken)
		if _, err := creds.Get(); err != nil {
			return errors.WithStack(err)
		}

		awsConf = awsConf.WithCredentials(creds)
	}

	svc := awss3.New(session.New(awsConf))
	input := &awss3.PutObjectInput{
		ACL:         aws.String(acl),
		Body:        aws.ReadSeekCloser(r),
		Bucket:      aws.String(bucket),
		ContentType: aws.String("application/pdf"),
		Key:         aws.String(key),
	}

	if _, err := svc.PutObject(input); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

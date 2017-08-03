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
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
func (*S3) Upload(ctx context.Context, r io.Reader, opts map[string]*proto.Option) error {
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

	sess, err := session.NewSession(awsConf)
	if err != nil {
		return errors.WithStack(err)
	}

	uploader := s3manager.NewUploader(sess)
	ct := "application/pdf"
	input := &s3manager.UploadInput{
		ACL:         &acl,
		Body:        r,
		Bucket:      &bucket,
		ContentType: &ct,
		Key:         &key,
	}

	if _, err := uploader.UploadWithContext(ctx, input, nil); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

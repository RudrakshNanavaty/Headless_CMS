package initializers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func CreateAWSSession() *s3manager.Uploader {
	awsBucketRegion := os.Getenv("AWS_BUCKET_REGION")
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecretKey := os.Getenv("AWS_SECRET")

	awsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(awsBucketRegion),
			Credentials: credentials.NewCredentials(&credentials.StaticProvider{
				Value: credentials.Value{
					AccessKeyID:     awsAccessKey,
					SecretAccessKey: awsSecretKey,
				},
			}),
		},
	})
	if err != nil {
		panic(err)
	}
	uploader := s3manager.NewUploader(awsSession)
	return uploader
}

var (
	Uploader = CreateAWSSession()
)

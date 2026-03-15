package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"mime/multipart"
	"os"
)

var uploader *s3manager.Uploader
var awsSession *session.Session

func SaveFile(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {
	if uploader == nil {
		var err error
		awsBucketRegion := os.Getenv("AWS_BUCKET_REGION")
		awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
		awsSecretKey := os.Getenv("AWS_SECRET")
		awsSession, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region: aws.String(awsBucketRegion),
				Credentials: credentials.NewStaticCredentials(
					awsAccessKey,
					awsSecretKey,
					"",
				),
			},
		})

		if err != nil {
			panic(err)
		}
	}

	uploader = s3manager.NewUploader(awsSession)
	// Upload the file to S3 using the fileReader
	awsBucketName := os.Getenv("AWS_BUCKET_NAME")
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awsBucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	// Get the URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", awsBucketName, fileHeader.Filename)

	return url, nil
}

package filesystem

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"sqs_s3_lambda/config"
)

var S3Client *s3.Client

func InitS3Client() {
	cfg := config.GetAwsCredentialConfig()
	S3Client = s3.NewFromConfig(cfg)
}

package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"os"
)

func GetAwsCredentialConfig() aws.Config {
	awsAccessKey := os.Getenv("ACCESS_KEY")
	awsSecretKey := os.Getenv("SECRET")
	region := os.Getenv("REGION")

	creds := credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")

	config := aws.Config{
		Region:      region,
		Credentials: creds,
	}

	return config
}

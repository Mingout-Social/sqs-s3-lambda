package main

import (
	"bytes"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
	"sqs_s3_lambda/config/filesystem"
	"sqs_s3_lambda/lib"
)

func handler(ctx context.Context, event events.SQSEvent) error {
	var err error

	filesystem.InitS3Client()

	for _, msg := range event.Records {
		reader := bytes.NewReader([]byte(msg.Body))

		_, err = lib.PutFile(msg.MessageId+".json", reader, reader.Size(), types.ObjectCannedACLPublicRead)
		log.Printf("Error in Uploading File: %v", err)
	}

	return err
}

func main() {
	lambda.Start(handler)
}

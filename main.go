package main

import (
	"bytes"
	"context"
	"github.com/Mingout-Social/mo-aws-lib/config/filesystem"
	"github.com/Mingout-Social/mo-aws-lib/lib"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
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

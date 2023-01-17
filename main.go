package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
	"sqs_s3_lambda/config/filesystem"
	"sqs_s3_lambda/lib"
)

type MessagePayload struct {
	MessageId   string `json:"message_id"`
	EventSource string `json:"event_source"`
	Body        string `json:"body"`
}

func handler(ctx context.Context, event events.SQSEvent) error {
	var err error

	filesystem.InitS3Client()

	for _, msg := range event.Records {
		payload := MessagePayload{
			MessageId:   msg.MessageId,
			EventSource: msg.EventSource,
			Body:        msg.Body,
		}

		file, _ := json.MarshalIndent(payload, "", "")

		reader := bytes.NewReader(file)

		_, err = lib.PutFile(payload.MessageId+".json", reader, reader.Size(), types.ObjectCannedACLPublicRead)
		log.Printf("Error in Uploading File: %v", err)
	}

	return err
}

func main() {
	lambda.Start(handler)
}

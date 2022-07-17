package sqsutil

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func PublishMessage(ctx context.Context, message *string, queueUrl *string) {

	client := GetSQSClient(ctx)
	client.SendMessage(ctx, nil)
	sendMessageInput := sqs.SendMessageInput{MessageBody: message, QueueUrl: queueUrl}
	fmt.Printf("Send message input is %+v", sendMessageInput)
}

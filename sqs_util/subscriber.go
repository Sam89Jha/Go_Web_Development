package sqsutil

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func RecieveMessage(ctx context.Context, queueUrl string) error {

	client := GetSQSClient(ctx)

	recieveMessageInput := &sqs.ReceiveMessageInput{
		QueueUrl:        &queueUrl,
		WaitTimeSeconds: 20,
	}
	fmt.Printf("Send message input is %+v\n", recieveMessageInput)
	for {
		resp, err := client.ReceiveMessage(ctx, recieveMessageInput)
		if err != nil {
			log.Fatal(err)
			return err
		}
		for _, v := range resp.Messages {
			fmt.Println(*v.Body)
			client.DeleteMessage(ctx, &sqs.DeleteMessageInput{QueueUrl: &queueUrl, ReceiptHandle: v.ReceiptHandle})
		}
	}
}

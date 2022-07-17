package sqsutil

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
)

func PublishMessage(ctx context.Context, message string, queueUrl string) (*sqs.SendMessageOutput, error) {

	client := GetSQSClient(ctx)

	var grpId string = "groupId1"
	uuidWithHyphen := uuid.New()
	fmt.Println(uuidWithHyphen)
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	var dedpulicationId string = uuid

	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:            &message,
		QueueUrl:               &queueUrl,
		MessageGroupId:         &grpId,
		MessageDeduplicationId: &dedpulicationId,
	}
	fmt.Printf("Send message input is %+v\n", sendMessageInput)

	resp, err := client.SendMessage(ctx, sendMessageInput)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp, nil
}

package sqsutil

import (
	"context"
	"log"
)

func GetSQSQueuesList(ctx context.Context) []string {
	queues := make([]string, 0)
	client := GetSQSClient(ctx)
	urls, err := client.ListQueues(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return queues
	}
	queues = append(queues, urls.QueueUrls...)
	return queues
}

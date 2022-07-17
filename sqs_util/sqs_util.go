package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	fmt.Print("Calling sqs client")

	client := GetSQSClient(ctx)
	output, err := client.ListQueues(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range output.QueueUrls {
		fmt.Printf("queue url is %s", v)
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	restapi "github.com/Sam89Jha/Go_Web_Development/rest_api"
	sqsutil "github.com/Sam89Jha/Go_Web_Development/sqs_util"
)

func main() {

	fmt.Println("Welcome to Rest API example !!")
	restapi.RegisterAPIs()

	fmt.Println("Welcome to AWS SQS practice")
	// ctx := context.Background()
	// subscribe(ctx)
	//publish(ctx)

}
func listSqsQueues(ctx context.Context) {
	queues := sqsutil.GetSQSQueuesList(ctx)
	fmt.Println(queues)
}
func publish(ctx context.Context) {
	fmt.Printf("Enter message to publish\n")
	var msg string
	fmt.Scanln(&msg)
	resp, err := sqsutil.PublishMessage(ctx, msg, sqsutil.DevQueueFifoUrl)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("response of publish message is %+v\n", resp)
}

func subscribe(ctx context.Context) {
	err := sqsutil.RecieveMessage(ctx, sqsutil.DevQueueFifoUrl)
	if err != nil {
		log.Fatal(err)
	}
}

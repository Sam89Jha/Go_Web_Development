package main

import (
	"context"
	"fmt"

	sqsutil "github.com/Sam89Jha/Go_Web_Development/sqs_util"
)

func main() {
	fmt.Println("AWS Queue listing")
	ctx := context.Background()
	fmt.Println(sqsutil.GetSQSQueuesList(ctx))
}

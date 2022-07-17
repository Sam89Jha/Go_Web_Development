package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func GetSQSClient(ctx context.Context) *sqs.Client {
	cfg := getAWSConfig(ctx)
	client := sqs.NewFromConfig(cfg)
	return client
}

func getAWSConfig(ctx context.Context) aws.Config {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

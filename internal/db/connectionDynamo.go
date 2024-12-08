package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/elfaldia/taller-noSQL/internal/env"
)

func ConnectDynamoDB() *dynamodb.Client {




	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(env.GetString("region", "sa-east-1")),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(env.GetString("aws_access_key_id", ""), env.GetString("aws_secret_access_key", ""), ""),
		),
	)
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	log.Print("You successfully connected to DynamoDB!")
	client := dynamodb.NewFromConfig(cfg)

	return client
}
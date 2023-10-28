package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DBService struct {
	Client *dynamodb.Client
}

func Connect() *DBService {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("ap-northeast-2"),
	)
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	return &DBService{Client: dynamodb.NewFromConfig(cfg)}
}

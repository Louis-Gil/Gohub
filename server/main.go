package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"server/handler"
	"server/infra/db"
)

var DB *db.DBService

func init() {
	DB = db.Connect()
}

func main() {
	lambda.Start(handler.HandleRequest)
}

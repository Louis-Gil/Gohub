package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"server/infra/db"
	"server/handler"
)

var DB *db.DBService

func init() {
	DB = db.Connect()
}

func main() {
	lambda.Start(handler.HandleRequest)
}

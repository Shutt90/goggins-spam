package main

import (
	"github.com/shutt90/goggins-spam/controllers"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func handleRequest() {
	godotenv.Load()

	controller.GetQuote()
}

func main() {
	lambda.Start(handleRequest)
}

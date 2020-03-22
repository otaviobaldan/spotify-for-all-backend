package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/otaviobaldan/spotify-for-all-backend/router"
)

func main() {
	lambda.Start(router.Router)
}

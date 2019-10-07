package main

import (
	github.com/aws/aws-lambda-go/lambda
)

func main() {
	lambda.start(validateReward)
}

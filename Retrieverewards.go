package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type rewardsRequest struct {
	Mobilenumber string `json:"Mobilenumber"`
}

type rewardsResponse struct {
	status string
}


var db = dynamodb.New(session.New(), aws.NewConfig())

func getRewards(rewardsReq rewardsRequest) {
	fmt.Println("rewardsReq in getRewards", rewardsReq.Mobilenumber)

	input := &dynamodb.GetItemInput{
		TableName: aws.String("ActiveRewards"),
		Key: map[string]*dynamodb.AttributeValue{
			"MobileNumber": {
				S: aws.String(rewardsReq.Mobilenumber),
			},
		},
	}
	fmt.Printf("Input data %v", input)
	result, err := db.GetItem(input)
	fmt.Printf("Result Data %v", result)

	if err != nil {
		fmt.Println("ERROR", err.Error())
		fmt.Println("Inside Check Rewards new code", rewardsReq.Mobilenumber, result)

	} else {
		fmt.Println("Inside Check Rewards Mobile Results", result)

	}
	rewardsResponse := new(rewardsResponse)
	rewardsResponse.status = "200"
	return

}

func main() {
	//var rewardsReq rewardsRequest
	lambda.Start(getRewards)
}

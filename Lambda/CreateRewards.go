package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//RewardEvent - Input object
type RewardEvent struct {
	MobileNumber      string `json:"MobileNumber"`
	CodeGeneratedDate string `json:"CodeGeneratedDate"`
	ConfirmationCode  int    `json:"ConfirmationCode"`
}

//HandleRequest - handles request
func HandleRequest(ctx context.Context, rewardEvent RewardEvent) (string, error) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	tableName := "ActiveRewards"

	/*rewardEvent = RewardEvent{
		MobileNumber:      "0432957930",
		CodeGeneratedDate: "1564723644",
		ConfirmationCode:  1234,
	}*/

	rewardEvent.CodeGeneratedDate = time.Now().String()

	av, err := dynamodbattribute.MarshalMap(rewardEvent)
	fmt.Println(av)
	if err != nil {
		fmt.Println("Got error marshalling new Rewards item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return "Successfully added item with confirmation code to table " + tableName, nil
}

func main() {
	lambda.Start(HandleRequest)
}

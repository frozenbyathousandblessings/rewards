package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	 "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

)

type rewardsRequest struct {
	Mobilenumber string `json:"Mobilenumber"`
}



type RewardsResponse struct {
    Status   string `json:"Status"`
	MobileNumber      string `json:"MobileNumber"`
	ConfirmationCode  int    `json:"ConfirmationCode"`
}

var db = dynamodb.New(session.New(), aws.NewConfig())

func getRewards(rewardsReq rewardsRequest) (RewardsResponse,error){
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
    rewardsResponse :=RewardsResponse{}
	if err != nil {
		fmt.Println("ERROR", err.Error())

	} else {

	    err = dynamodbattribute.UnmarshalMap(result.Item, &rewardsResponse)
        if(err!=nil){
            fmt.Println("Error During Unmarshal")
         }
         if(len(rewardsResponse.MobileNumber)!=0){
            rewardsResponse.Status="200"
         }else {
              rewardsResponse.Status="404"
         }
      }

    fmt.Printf("RewardsResponse {}",rewardsResponse)
	return rewardsResponse,nil

}

func main() {
	//var rewardsReq rewardsRequest
	lambda.Start(getRewards)
}

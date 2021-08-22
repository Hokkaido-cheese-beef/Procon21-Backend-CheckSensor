package dao

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type  DynamoDB struct{
	Dynamo  *dynamodb.DynamoDB
	CheckDevice Methods
}

type Methods struct {
	CheckDeviceLogic methods
}

type methods interface {
	CheckDeviceExist(deviceID string)error
	CheckDeviceMotion(deviceID string)(int,error)
}

func New()(*DynamoDB,error){
	//DB接続
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// init methods
	checkDeviceMethod := newCheckDeviceClient(svc)

	return &DynamoDB{
		Dynamo:  svc,
		CheckDevice: Methods{checkDeviceMethod},
	},nil
}

package dao

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"login/pkg/model/dto"
)

type  DynamoDB struct{
	Dynamo  *dynamodb.DynamoDB
	Login Methods
}

type Methods struct {
	LoginLogic methods
}

type methods interface {
	GetLoginUser(req dto.LoginReq)error
}

func New()(*DynamoDB,error){
	//DB接続
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// init methods
	loginMethod := newLoginClient(svc)

	return &DynamoDB{
		Dynamo:  svc,
		Login: Methods{loginMethod},
	},nil
}

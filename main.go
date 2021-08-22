package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"login/pkg/model/dao"
	"login/pkg/model/dto"
	"login/pkg/res"
)



func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse,error) {
	var req dto.LoginReq
	var response dto.Response
	err :=json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		log.Println(err)
		return res.ReturnInternalServerErrorResponse(err)
	}

	client, err := dao.New()
	if err != nil {
		log.Println(err)
		return res.ReturnInternalServerErrorResponse(err)
	}

	err = client.Login.LoginLogic.GetLoginUser(req)
	if err != nil {
		if err.Error()=="user is not exits" {
			response.ErrorMessage="user is not exits"
			responseBody, _ := json.Marshal(response)
			return res.ReturnBadRequestResponse(string(responseBody)), nil
		}else if err.Error()=="password is wrong"{
			response.ErrorMessage="password is wrong"
			responseBody, _ := json.Marshal(response)
			return res.ReturnBadRequestResponse(string(responseBody)), nil
		}
		return res.ReturnInternalServerErrorResponse(err)
	}

	responseBody, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       string(responseBody),
		StatusCode: 200,
	},nil
}

func main(){
	lambda.Start(handler)
}

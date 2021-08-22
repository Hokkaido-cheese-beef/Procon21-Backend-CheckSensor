package res

import "github.com/aws/aws-lambda-go/events"

func ReturnBadRequestResponse(errMessage string)events.APIGatewayProxyResponse{
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       errMessage,
	}
}

func ReturnInternalServerErrorResponse(err error)(events.APIGatewayProxyResponse,error){
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	},err
}


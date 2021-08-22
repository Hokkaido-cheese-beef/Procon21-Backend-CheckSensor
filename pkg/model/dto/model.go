package dto

type LoginReq struct{
	UserID string `dynamodbav:"userID" json:"userID"`
	Password string `dynamodbav:"password" json:"password"`
}


type User struct{
	UserID string `dynamodbav:"userID"`
	Created int `dynamodbav:"createdAt"`
	Password string `dynamodbav:"password"`
}

type Response struct{
	ErrorMessage string `json:"errorMessage"`
}
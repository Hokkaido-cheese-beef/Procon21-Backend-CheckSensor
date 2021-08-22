package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"login/pkg/model/dto"
	"errors"
	"strconv"
	"time"
)

type loginMethods struct {
	Client *dynamodb.DynamoDB
}

func newLoginClient(client *dynamodb.DynamoDB) methods {
	return &loginMethods{Client: client}
}

func (r *loginMethods) 	GetLoginUser(req dto.LoginReq)error {
	var err error
	now :=time.Now()
	input := &dynamodb.QueryInput{
		TableName: aws.String("user"),
		ExpressionAttributeNames: map[string]*string{
			"#userID":   aws.String("userID"), // alias付けれたりする
			"#createdAt": aws.String("createdAt"),   // 予約語はそのままだと怒られるので置換する
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userID": { // :を付けるのがセオリーのようです
				S: aws.String(req.UserID),
			},
			":timestamp": { // :を付けるのがセオリーのようです
				N: aws.String(strconv.FormatInt(now.Unix(), 10)),
			},
		},
		KeyConditionExpression: aws.String("#userID = :userID AND #createdAt < :timestamp"),         // 検索条件
		ScanIndexForward:       aws.Bool(false),                 // ソートキーのソート順（指定しないと昇順）
		Limit:                  aws.Int64(1),                  // 取得件数の指定もできる
	}

	results, err := r.Client.Query(input)
	if err != nil {
		log.Println(err)
		return 	err
	}

	user := &dto.User{}
	for _,result :=range results.Items{
		err = dynamodbattribute.UnmarshalMap(result, user)
		if err != nil {
			fmt.Println("[Unmarshal Error]", err)
			return err
		}
	}

	if user.UserID==""{
		return errors.New("user is not exits")
	}

	if user.Password!=req.Password{
		return errors.New("password is wrong")
	}

	return nil
}

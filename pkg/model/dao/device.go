package dao

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"checkDevice/pkg/model/dto"
	"errors"
	"strconv"
	"time"
)

type checkDeviceClientMethods struct {
	Client *dynamodb.DynamoDB
}

func newCheckDeviceClient(client *dynamodb.DynamoDB) methods {
	return &checkDeviceClientMethods{Client: client}
}

func (r *checkDeviceClientMethods) 	CheckDeviceExist(deviceID string)error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("deviceInfo"),
		Key: map[string]*dynamodb.AttributeValue{
			"deviceID": {
				N: aws.String(deviceID),
			},
		},
	}

	result, err := r.Client.GetItem(input)
	if err != nil {
		log.Println("[GetItem Error]", err)
		return err
	}

	deviceInfo := &dto.DeviceInfo{}
	if err := dynamodbattribute.UnmarshalMap(result.Item, deviceInfo); err != nil {
		log.Println("[Unmarshal Error]", err)
		return err
	}
	if deviceInfo.DeviceID==""{
		return errors.New("deviceID is wrong")
	}
	return nil
}

func (r *checkDeviceClientMethods)CheckDeviceMotion(deviceID string)(int,error){
	now :=time.Now()
	beforeTime:=strconv.FormatInt(now.Unix()-600, 10)

	input := &dynamodb.QueryInput{
		TableName: aws.String("SensorData"),
		ExpressionAttributeNames: map[string]*string{
			"#sensorID":   aws.String("sensorID"), // alias付けれたりする
			"#timestamp": aws.String("timestamp"),   // 予約語はそのままだと怒られるので置換する
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":sensorID": { // :を付けるのがセオリーのようです
				S: aws.String(deviceID),
			},
			":timestampbefore": { // :を付けるのがセオリーのようです
				N: aws.String(beforeTime),
			},
			":timestampafter": { // :を付けるのがセオリーのようです
				N: aws.String(strconv.FormatInt(now.Unix(), 10)),
			},
		},
		KeyConditionExpression: aws.String("#sensorID = :sensorID AND #timestamp BETWEEN :timestampbefore AND :timestampafter"),         // 検索条件
		ScanIndexForward:       aws.Bool(false),                 // ソートキーのソート順（指定しないと昇順）
		Limit:                  aws.Int64(1),                  // 取得件数の指定もできる
	}

	results, err := r.Client.Query(input)
	if err != nil {
		log.Println(err)
		return 	0,err
	}

	sensorData := &dto.SensorData{}
	for _,result :=range results.Items{
		err = dynamodbattribute.UnmarshalMap(result, sensorData)
		if err != nil {
			log.Println("[Unmarshal Error]", err)
			return 0,err
		}
	}

	if sensorData.Co2==0 && sensorData.Hum==0 && sensorData.Temp==0{
		return 0,err
	}
	return 1,nil
}

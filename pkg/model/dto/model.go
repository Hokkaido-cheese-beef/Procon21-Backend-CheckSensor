package dto

type DeviceInfo struct{
	DeviceID string `dynamodbav:"deviceID"`
}

type SensorData  struct{
	SensorID string `dynamodbav:"sensorID" json:"sensorID"`
	Timestamp int `dynamodbav:"timestamp" json:"timestamp"`
	Co2 int `dynamodbav:"co2" json:"co2"`
	Hum float64 `dynamodbav:"hum" json:"hum"`
	Temp float64 `dynamodbav:"temp" json:"temp"`
}

type Response struct{
	Message string `json:"Message"`
}
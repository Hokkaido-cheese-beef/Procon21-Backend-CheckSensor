# Procon2021CheckDevice

## 基本仕様
---
### 概要
デバイスの動作確認を行います。

### 特徴
リクエストで渡されたデバイスが動作しているかどうか確認します。

## 機能仕様
- HTTP method:GET
- endpoint:/sensor/{ここにdeviceID}

---
### 利用法

- endpoint:/sensor/{ここにdeviceID}

####  resqonse
```cassandraql
{
  "Message": "action"
}
or
{
  "Message": "not action"
}
```
他にもIDが正しくない場合もメッセージを返却します。

---
### アーキテクチャ
MVCもどき

### 依存環境
- AWS (Lambda)
- AWS (API Gateway)

### 使用ライブラリ
主要なライブラリのみを示す
- aws-sdk()
- Gin(Webフレームワーク)

### デプロイ


### セットアップ
対象となるソースコードをzipに圧縮してアップロードする必要があります
2. GOOS=linux GOARCH=amd64 go build -o hello main.go
3. zip function.zip hello

### 注意点
動作確認はAWSのLambdaにアップロードする必要があります
zipはGitにあげないこと

### 作成者
Taketo Wakamatsu (若松丈人)

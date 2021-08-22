# Procon21-Backend-Login

## 基本仕様
---
### 概要
ユーザーのログイン認証を行います。

### 特徴
今回の使用は特にパスワードを暗号化する、JWTを使用した使用にするｍなどといったことはしていません。
ユーザーが存在するか、パスワードが違っていないかを確認し認証します。

## 機能仕様
- HTTP method:POST
- endpoint:/login

---
### 利用法

- endpoint:/login

#### request
```cassandraql
{
    "userID":"0001",
    "password":"0001"
}
```

####  resqonse
```cassandraql
{
  "errorMessage": ""
}
```

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

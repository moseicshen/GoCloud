## Init
```
goctl api new core
goctl api go -api core.api -dir . -style go_zero
go run core.go -f etc/core-api.yaml
```

## Tools
Xorm: do not forget to add mysql driver
```
_ "github.com/go-sql-driver/mysql"
```
Email code sender:
```
https://github.com/jordan-wright/email
```
Go Redis
```
https://github.com/redis/go-redis
---
# start server
redis-server

# connect to redis
redis-cli
```
Go UUID
```
go get github.com/google/uuid
```
Tencent COS
```
platform: https://console.cloud.tencent.com/cos/bucket
document: https://cloud.tencent.com/document/product/436/31215
```
Go-zero Middleware
```
https://go-zero.dev/docs/tutorials/api/middleware
```
```
goctl api new core
goctl api go -api core.api -dir . -style go_zero
go run core.go -f etc/core-api.yaml
```

```
xorm: do not forget to add mysql driver
_ "github.com/go-sql-driver/mysql"

email code sender:
https://github.com/jordan-wright/email
```
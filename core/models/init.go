package models

import (
	"GoCloud/core/key"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()
var RedisDB = InitRedis()

func Init() *xorm.Engine {
	sourceStr := fmt.Sprintf("%s:%s@/GoCloud?charset=utf8", key.DBUser, key.DBPwd)
	engine, err := xorm.NewEngine("mysql", sourceStr)
	if err != nil {
		log.Printf("Create New Xorm DB Failure %v\n", err)
	}
	return engine
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

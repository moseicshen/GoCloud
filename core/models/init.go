package models

import (
	"GoCloud/core/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"log"
	"xorm.io/xorm"
)

func InitMysql(c config.Config) *xorm.Engine {
	dataSource := c.Mysql.DataSource
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Create New Xorm DB Failure %v\n", err)
	}
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

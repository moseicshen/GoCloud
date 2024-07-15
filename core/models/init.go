package models

import (
	"GoCloud/core/key"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	sourceStr := fmt.Sprintf("%s:%s@/GoCloud?charset=utf8", key.DBUser, key.DBPwd)
	engine, err := xorm.NewEngine("mysql", sourceStr)
	if err != nil {
		log.Printf("Create New Xorm DB Failure %v\n", err)
	}
	return engine
}

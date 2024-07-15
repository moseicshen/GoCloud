package test

import (
	"GoCloud/core/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func TestXorm(t *testing.T) {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:1234@/GoCloud?charset=utf8")
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.UserBasic, 0)
	err = engine.Find(data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
	for _, item := range data {
		fmt.Println(item)
	}
}

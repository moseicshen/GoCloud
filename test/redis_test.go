package test

import (
	"GoCloud/core/models"
	"context"
	"log"
	"testing"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	result, err := models.RedisDB.Get(ctx, "zxshenmonica@163.com").Result()
	if err != nil {
		return
	}
	log.Printf(result)
}

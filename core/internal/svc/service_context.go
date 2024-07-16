package svc

import (
	"GoCloud/core/internal/config"
	"GoCloud/core/models"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config  config.Config
	Engine  *xorm.Engine
	RedisDB *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Engine:  models.InitMysql(c),
		RedisDB: models.InitRedis(c),
	}
}

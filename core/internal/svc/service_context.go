package svc

import (
	"GoCloud/core/internal/config"
	"GoCloud/core/internal/middleware"
	"GoCloud/core/models"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config  config.Config
	Engine  *xorm.Engine
	RedisDB *redis.Client
	Auth    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Engine:  models.InitMysql(c),
		RedisDB: models.InitRedis(c),
		Auth:    middleware.NewAuthMiddleware().Handle,
	}
}

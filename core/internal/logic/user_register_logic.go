package logic

import (
	"GoCloud/core/helper"
	"GoCloud/core/models"
	"context"
	"errors"
	"log"
	"strconv"

	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// compare the input code with the one stored in redis
	code, err := models.RedisDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("request for email code first")
	}
	if code != req.Code {
		err = errors.New("wrong email code")
		return nil, err
	}

	// the username was not registered before
	cnt, err := models.Engine.Where("name = ?", req.UserName).Count(&models.UserBasic{})
	if err != nil {
		return nil, err
	}
	if cnt != 0 {
		err = errors.New("username already registered")
		return nil, err
	}

	//insert user info into database
	uuid := helper.UUID()
	if uuid == "" {
		return nil, errors.New("uuid generation error")
	}
	user := models.UserBasic{
		Identity: uuid,
		Name:     req.UserName,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	n, err := models.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	models.RedisDB.Del(l.ctx, req.Email)
	log.Printf("insert user row: %s", strconv.FormatInt(n, 10))
	return
}

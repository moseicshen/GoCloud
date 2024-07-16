package logic

import (
	"GoCloud/core/define"
	"GoCloud/core/helper"
	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"
	"GoCloud/core/models"
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeLogic {
	return &EmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCodeLogic) EmailCode(req *types.EmailCodeRequest) (resp *types.EmailCodeResponse, err error) {
	// the email was not registered before
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("the Email has been registered")
		return
	}
	// generate a random code
	code := helper.RandCode()
	// store code in redis
	l.svcCtx.RedisDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.EmailCodeExpireTime))
	// send code to email
	err = helper.SendEmailCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}

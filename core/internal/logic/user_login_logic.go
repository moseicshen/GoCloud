package logic

import (
	"GoCloud/core/helper"
	"GoCloud/core/models"
	"context"
	"errors"

	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// Find user from user database
	user := new(models.UserBasic)
	get, err := models.Engine.Where("name = ? AND password = ?", req.UserName, helper.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("wrong User or Password")
	}
	// Get token from token service
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token
	return
}

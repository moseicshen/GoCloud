package logic

import (
	"GoCloud/core/models"
	"context"
	"errors"

	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	resp = new(types.InfoResponse)
	ub := new(models.UserBasic)
	get, err := models.Engine.Where("identity=?", req.Identity).Get(ub)
	if err != nil {
		return nil, err
	}
	if !get {
		return nil, errors.New("wrong User Identity")
	}
	resp.UserName = ub.Name
	resp.Email = ub.Email
	return
}

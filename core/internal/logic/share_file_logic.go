package logic

import (
	"GoCloud/core/helper"
	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"
	"GoCloud/core/models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileLogic {
	return &ShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileLogic) ShareFile(req *types.ShareFileRequest, userIdentity string) (resp *types.ShareFileResponse, err error) {
	ur := new(models.UserRepository)
	_, err = l.svcCtx.Engine.Where("identity = ?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return nil, err
	}

	// insert share basic data
	data := models.ShareBasic{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		RepositoryIdentity: ur.RepositoryIdentity,
		FileName:           ur.Name,
		ExpiredTime:        req.ExpiredTime,
		ClickCount:         0,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	resp = new(types.ShareFileResponse)
	resp.Identity = data.Identity
	code := helper.RandCode()
	resp.ShareCode = code
	return
}

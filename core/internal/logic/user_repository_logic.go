package logic

import (
	"GoCloud/core/helper"
	"GoCloud/core/models"
	"context"

	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositoryLogic {
	return &UserRepositoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositoryLogic) UserRepository(req *types.UserRepositoryRequest, userIdentity string) (resp *types.UserRepositoryResponse, err error) {
	ur := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	resp = new(types.UserRepositoryResponse)
	resp.Identity = ur.Identity
	return

}

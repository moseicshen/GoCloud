package logic

import (
	"GoCloud/core/helper"
	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"
	"GoCloud/core/models"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
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
	for {
		count, err := l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ? AND name = ?", req.ParentId, userIdentity, req.Name).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			prevName, _ := strings.CutSuffix(req.Name, req.Ext)
			req.Name = fmt.Sprintf("%s_1%s", prevName, req.Ext)
		} else {
			break
		}
	}
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

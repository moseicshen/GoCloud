package logic

import (
	"GoCloud/core/helper"
	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"
	"GoCloud/core/models"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type FolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FolderCreateLogic {
	return &FolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FolderCreateLogic) FolderCreate(req *types.FolderCreateRequest, userIdentity string) (resp *types.FolderCreateResponse, err error) {
	for {
		count, err := l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ? AND name = ?", req.ParentId, userIdentity, req.Name).Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			req.Name = fmt.Sprintf("%s_1", req.Name)
		} else {
			break
		}
	}
	ur := &models.UserRepository{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	resp = new(types.FolderCreateResponse)
	resp.Identity = ur.Identity
	return
}

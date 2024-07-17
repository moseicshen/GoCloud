package logic

import (
	"GoCloud/core/models"
	"context"

	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileRenameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileRenameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileRenameLogic {
	return &FileRenameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileRenameLogic) FileRename(req *types.FileRenameRequest, userIdentity string) (resp *types.FileRenameResponse, err error) {
	file := &models.UserRepository{
		Name: req.NewName,
	}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(file)
	if err != nil {
		return nil, err
	}
	return
}

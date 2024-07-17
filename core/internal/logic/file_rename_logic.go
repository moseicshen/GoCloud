package logic

import (
	"GoCloud/core/models"
	"context"
	"errors"

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
	// whether the newName exists in the dir
	ub := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Get(ub)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("file not found")
	}
	count, err := l.svcCtx.Engine.Table("user_repository").Where("name = ? AND parent_id = ?", req.NewName, ub.ParentId).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("file name already exists")
	}

	// insert file with new name
	file := &models.UserRepository{
		Name: req.NewName,
	}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(file)
	if err != nil {
		return nil, err
	}
	return
}

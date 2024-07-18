package logic

import (
	"GoCloud/core/models"
	"context"
	"errors"

	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveFileLogic {
	return &MoveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveFileLogic) MoveFile(req *types.MoveFileRequest, userIdentity string) (resp *types.MoveFileResponse, err error) {
	ur := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).Get(ur)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("file not found")
	}
	updateRecord := &models.UserRepository{
		ParentId: ur.Id,
	}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(updateRecord)
	if err != nil {
		return nil, err
	}
	return
}

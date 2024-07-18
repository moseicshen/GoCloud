package logic

import (
	"GoCloud/core/models"
	"context"

	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSharedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckSharedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSharedLogic {
	return &CheckSharedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckSharedLogic) CheckShared(req *types.CheckSharedRequest) (resp *types.CheckSharedResponse, err error) {
	// update share_basic click_count
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Incr("click_count").Update(new(models.ShareBasic))
	if err != nil {
		return nil, err
	}

	// get file info
	resp = new(types.CheckSharedResponse)
	_, err = l.svcCtx.Engine.Table("share_basic").Where("share_basic.identity = ?", req.Identity).
		Select("share_basic.repository_identity, share_basic.file_name, repository_pool.ext, repository_pool.size, repository_pool.path").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Get(resp)
	if err != nil {
		return nil, err
	}
	return
}

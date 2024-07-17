package logic

import (
	"GoCloud/core/define"
	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositoryListLogic {
	return &UserRepositoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositoryListLogic) UserRepositoryList(req *types.UserRepositoryListRequest, userIdentity string) (resp *types.UserRepositoryListResponse, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserRepositoryListResponse)

	size := req.Size
	if size <= 0 {
		size = define.ListPageSizeDefault
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageOffset := (page - 1) * size

	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.name, user_repository.ext, repository_pool.size, repository_pool.path").
		Join("INNER", "repository_pool", "repository_pool.identity = user_repository.repository_identity").
		Limit(size, pageOffset).Find(&uf)
	if err != nil {
		return nil, err
	}
	count, err := l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count()
	if err != nil {
		return nil, err
	}

	if pageOffset > int(count) {
		return nil, errors.New("page out of range")
	}

	resp.List = uf
	resp.Count = int(count)
	return
}

package handler

import (
	"net/http"

	"GoCloud/core/internal/logic"
	"GoCloud/core/internal/svc"
	"GoCloud/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MoveFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MoveFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMoveFileLogic(r.Context(), svcCtx)
		resp, err := l.MoveFile(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

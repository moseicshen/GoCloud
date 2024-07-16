// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"GoCloud/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/email/sendcode",
				Handler: EmailCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload",
				Handler: FileUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
		},
	)
}

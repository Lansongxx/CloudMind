// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	login "CloudMind/internal/handler/login"
	tools "CloudMind/internal/handler/tools"
	"CloudMind/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login/Register",
				Handler: login.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/EmailLogin",
				Handler: login.EmailLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/QqLogin",
				Handler: login.QqloginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/WxLogin",
				Handler: login.WxloginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/tools/SendEmail",
				Handler: tools.SendEmailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)
}
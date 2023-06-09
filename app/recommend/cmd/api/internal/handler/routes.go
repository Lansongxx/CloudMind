// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	recommend "CloudMind/app/recommend/cmd/api/internal/handler/recommend"
	"CloudMind/app/recommend/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/GetRecommendByUserId/:number",
				Handler: recommend.GetRecommendByUserIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/GetRecommendByItemId/:itemId/:number",
				Handler: recommend.GetRecommendByItemIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/InsertFeedBack",
				Handler: recommend.InsertFeedBackHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/recommend/v1"),
	)
}

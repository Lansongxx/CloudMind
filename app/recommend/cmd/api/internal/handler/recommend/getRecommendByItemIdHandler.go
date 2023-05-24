package recommend

import (
	"net/http"

	"CloudMind/app/recommend/cmd/api/internal/logic/recommend"
	"CloudMind/app/recommend/cmd/api/internal/svc"
	"CloudMind/app/recommend/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRecommendByItemIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRecommendByItemIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := recommend.NewGetRecommendByItemIdLogic(r.Context(), svcCtx)
		resp, err := l.GetRecommendByItemId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

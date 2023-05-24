package recommend

import (
	"CloudMind/app/recommend/cmd/rpc/pb"
	"context"

	"CloudMind/app/recommend/cmd/api/internal/svc"
	"CloudMind/app/recommend/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendByItemIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecommendByItemIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendByItemIdLogic {
	return &GetRecommendByItemIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecommendByItemIdLogic) GetRecommendByItemId(req *types.GetRecommendByItemIdReq) (*types.GetRecommendByItemIdResp, error) {
	Resp, err := l.svcCtx.RecommendRpc.GetRecommendByItem(l.ctx, &pb.GetRecommendByItemIdReq{
		ItemId: req.ItemId,
		Number: req.Number,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetRecommendByItemIdResp{
		ItemIds: Resp.ItemId,
	}, nil
}

package logic

import (
	"context"

	"CloudMind/app/recommend/cmd/rpc/internal/svc"
	"CloudMind/app/recommend/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendByItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecommendByItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendByItemLogic {
	return &GetRecommendByItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecommendByItemLogic) GetRecommendByItem(in *pb.GetRecommendByItemIdReq) (*pb.GetRecommendByItemIdResp, error) {
	Resp, err := l.svcCtx.Gorse.GetNeighbors(context.Background(), in.ItemId, int(in.Number))
	if err != nil {
		return nil, err
	}
	var ItemIds []string
	for _, ItemId := range Resp {
		ItemIds = append(ItemIds, ItemId.Id)
	}

	return &pb.GetRecommendByItemIdResp{
		ItemId: ItemIds,
		Error:  "",
	}, nil
}

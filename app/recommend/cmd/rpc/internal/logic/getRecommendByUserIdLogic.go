package logic

import (
	"context"

	"CloudMind/app/recommend/cmd/rpc/internal/svc"
	"CloudMind/app/recommend/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRecommendByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendByUserIdLogic {
	return &GetRecommendByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRecommendByUserIdLogic) GetRecommendByUserId(in *pb.GetRecommendByUserIdReq) (*pb.GetRecommendByUserIdResp, error) {
	Resp, err := l.svcCtx.Gorse.GetRecommend(context.Background(), in.UserId, "", int(in.Number))
	if err != nil {
		return nil, err
	}

	return &pb.GetRecommendByUserIdResp{
		ItemId: Resp,
		Error:  "",
	}, nil
}

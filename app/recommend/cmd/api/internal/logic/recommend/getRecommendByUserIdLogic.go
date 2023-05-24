package recommend

import (
	"CloudMind/app/recommend/cmd/rpc/pb"
	"CloudMind/common/ctxdata"
	"context"
	"strconv"

	"CloudMind/app/recommend/cmd/api/internal/svc"
	"CloudMind/app/recommend/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRecommendByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecommendByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRecommendByUserIdLogic {
	return &GetRecommendByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecommendByUserIdLogic) GetRecommendByUserId(req *types.GetRecommendByUserIdReq) (resp *types.GetRecommendByUserIdResp, err error) {
	UserId := ctxdata.GetUidFromCtx(l.ctx)
	Resp, err := l.svcCtx.RecommendRpc.GetRecommendByUserId(l.ctx, &pb.GetRecommendByUserIdReq{
		UserId: strconv.Itoa(int(UserId)),
		Number: req.Number,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetRecommendByUserIdResp{
		ItemIds: Resp.ItemId,
	}, nil
}

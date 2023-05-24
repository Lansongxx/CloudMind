package recommend

import (
	"CloudMind/app/recommend/cmd/rpc/pb"
	"context"

	"CloudMind/app/recommend/cmd/api/internal/svc"
	"CloudMind/app/recommend/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertFeedBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertFeedBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertFeedBackLogic {
	return &InsertFeedBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertFeedBackLogic) InsertFeedBack(req *types.InsertFeedBackReq) (*types.InsertFeedBackResp, error) {
	var FeedBacks []*pb.FeedBack
	for _, FeedBack := range req.FeedBacks {
		FeedBacks = append(FeedBacks, &pb.FeedBack{
			FeedbackType: FeedBack.FeedbackType,
			UserId:       FeedBack.UserId,
			ItemId:       FeedBack.ItemId,
			Timestamp:    FeedBack.Timestamp,
		})
	}
	_, err := l.svcCtx.RecommendRpc.InsertFeedBack(l.ctx, &pb.InsertFeedBackReq{
		Feedback: FeedBacks,
	})
	if err != nil {
		return nil, err
	}
	return &types.InsertFeedBackResp{}, nil
}

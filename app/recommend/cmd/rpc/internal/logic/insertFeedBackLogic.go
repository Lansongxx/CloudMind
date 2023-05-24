package logic

import (
	"CloudMind/app/recommend/cmd/rpc/internal/svc"
	"CloudMind/app/recommend/cmd/rpc/pb"
	"context"
	"github.com/zhenghaoz/gorse/client"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertFeedBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertFeedBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertFeedBackLogic {
	return &InsertFeedBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertFeedBackLogic) InsertFeedBack(in *pb.InsertFeedBackReq) (*pb.InsertFeedBackResp, error) {
	var feedbacks []client.Feedback
	for _, feedback := range in.Feedback {
		feedbacks = append(feedbacks, client.Feedback{
			FeedbackType: feedback.FeedbackType,
			UserId:       feedback.UserId,
			ItemId:       feedback.ItemId,
			Timestamp:    feedback.Timestamp,
		})
	}
	_, err := l.svcCtx.Gorse.InsertFeedback(context.Background(), feedbacks)
	if err != nil {
		return nil, err
	}
	return &pb.InsertFeedBackResp{}, nil
}

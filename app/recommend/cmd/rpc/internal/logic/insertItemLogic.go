package logic

import (
	"context"
	"github.com/zhenghaoz/gorse/client"

	"CloudMind/app/recommend/cmd/rpc/internal/svc"
	"CloudMind/app/recommend/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertItemLogic {
	return &InsertItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertItemLogic) InsertItem(in *pb.InsertItemReq) (*pb.InsertItemResp, error) {
	_, err := l.svcCtx.Gorse.InsertItem(context.Background(), client.Item{
		ItemId:     in.Item.ItemId,
		IsHidden:   in.Item.IsHidden,
		Labels:     in.Item.Labels,
		Categories: in.Item.Categories,
		Timestamp:  in.Item.Timestamp,
		Comment:    in.Item.Comment,
	})
	if err != nil {
		return nil, err
	}
	return &pb.InsertItemResp{}, nil
}

package logic

import (
	"CloudMind/app/recommend/cmd/rpc/internal/svc"
	"CloudMind/app/recommend/cmd/rpc/pb"
	"context"
	"github.com/zhenghaoz/gorse/client"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemLogic {
	return &UpdateItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateItemLogic) UpdateItem(in *pb.UpdateItemReq) (*pb.UpdateItemResp, error) {
	times, err := time.Parse(in.ItemPatch.Timestamp, "2006-01-02 15:04:05")
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Gorse.UpdateItem(context.Background(), in.ItemId, client.ItemPatch{
		IsHidden:   &in.ItemPatch.IsHidden,
		Categories: in.ItemPatch.Categories,
		Timestamp:  &times,
		Labels:     in.ItemPatch.Labels,
		Comment:    &in.ItemPatch.Comment,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateItemResp{}, nil
}

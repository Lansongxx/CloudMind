package logic

import (
	"context"

	"CloudMind/app/recommend/cmd/rpc/internal/svc"
	"CloudMind/app/recommend/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteItemLogic {
	return &DeleteItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteItemLogic) DeleteItem(in *pb.DeleteItemReq) (*pb.DeleteItemResp, error) {
	_, err := l.svcCtx.Gorse.DeleteItem(context.Background(), in.ItemId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteItemResp{}, nil
}

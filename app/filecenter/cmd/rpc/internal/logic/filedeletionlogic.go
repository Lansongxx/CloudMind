package logic

import (
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeletionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDeletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeletionLogic {
	return &FileDeletionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDeletionLogic) FileDeletion(in *pb.FileDeletionReq) (*pb.FileDeletionResp, error) {

	_, err := l.svcCtx.FileModel.TxDeletes(l.ctx, l.svcCtx.GormDB, in.List)

	if err != nil {
		return nil, err
	}

	return &pb.FileDeletionResp{}, nil
}

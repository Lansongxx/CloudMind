package logic

import (
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMoveLogic {
	return &FileMoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileMoveLogic) FileMove(in *pb.FileMoveReq) (*pb.FileMoveResp, error) {

	// 手动开事务 在里面分别调用不同表格
	l.svcCtx.FileModel.TxDelete(l.ctx, l.svcCtx.GormDB, in.Id)

	return &pb.FileMoveResp{}, nil
}

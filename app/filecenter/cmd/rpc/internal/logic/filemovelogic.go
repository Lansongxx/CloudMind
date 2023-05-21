package logic

import (
	"CloudMind/app/filecenter/model"
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

	resp, err := l.svcCtx.FileModel.UpdateOneMapById(l.ctx, in.Id, map[string]interface{}{
		"parent_id": in.PreParentId,
	})

	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	if resp == 0 {
		return &pb.FileMoveResp{
			Error: "没有该文件",
		}, nil
	}

	return &pb.FileMoveResp{
		Error: "移动成功",
	}, nil
}

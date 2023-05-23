package file

import (
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilemoveLogic {
	return &FilemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilemoveLogic) Filemove(req *types.FileMoveReq) (*types.FileMoveResp, error) {

	resp, err := l.svcCtx.FileRpc.FileMove(l.ctx, &pb.FileMoveReq{
		LastParentId: req.LastParentId,
		PreParentId:  req.PreParentId,
		Id:           req.Id,
	})

	if err != nil {
		return nil, err
	}

	if resp.Error != "移动成功" {
		return nil, errorx.NewDefaultError(resp.Error)
	}

	return nil, nil
}

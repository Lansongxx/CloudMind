package file

import (
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FiledeletionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFiledeletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FiledeletionLogic {
	return &FiledeletionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FiledeletionLogic) Filedeletion(req *types.FileDeletionReq) (*types.FileDeletionResp, error) {

	var ids, parentids []int64
	for _, prefix := range req.Delist {
		parentids = append(parentids, prefix.ParentId)
		ids = append(ids, prefix.Id)
	}

	resp, err := l.svcCtx.FileRpc.FileDeletion(l.ctx, &pb.FileDeletionReq{
		ParentId: parentids,
		Id:       ids,
	})

	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errorx.NewDefaultError(resp.Error)
	}

	return nil, nil
}

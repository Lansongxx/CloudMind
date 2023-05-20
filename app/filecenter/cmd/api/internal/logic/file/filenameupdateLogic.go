package file

import (
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilenameupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilenameupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilenameupdateLogic {
	return &FilenameupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilenameupdateLogic) Filenameupdate(req *types.FileNameUpdateReq) (*types.FileNameUpdateResp, error) {

	resp, err := l.svcCtx.FileRpc.FileNameUpdate(l.ctx, &pb.FileNameUpdateReq{
		Id:   req.Id,
		Name: req.Name,
	})

	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errorx.NewDefaultError(resp.Error)
	}

	return nil, nil
}

package file

import (
	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"
	"CloudMind/app/filecenter/cmd/rpc/filecenter"
	"CloudMind/common/errorx"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilecreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilecreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilecreateLogic {
	return &FilecreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilecreateLogic) Filecreate(req *types.FileCreateReq) (*types.FileCreateResp, error) {

	resp, err := l.svcCtx.FileRpc.FileCreate(l.ctx, &filecenter.FileCreateReq{
		ParentId: req.ParentId,
		Name:     req.Name,
	})

	if err != nil {
		return nil, err
	}

	if resp.Error == "创建失败" {
		return nil, errorx.NewDefaultError("创建失败") // 使用 errorx 报错
	}

	return nil, nil
}

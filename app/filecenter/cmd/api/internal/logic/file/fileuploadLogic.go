package file

import (
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileuploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileuploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileuploadLogic {
	return &FileuploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileuploadLogic) Fileupload(req *types.FileUploadReq) (*types.FileUploadResp, error) {

	resp, err := l.svcCtx.FileRpc.FileUpload(l.ctx, &pb.FileUploadReq{
		Name:       req.Name,
		Type:       req.Type,
		SourcePath: req.SourcePath,
	})

	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errorx.NewDefaultError(resp.Error)
	}

	return nil, nil
}

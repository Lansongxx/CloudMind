package file

import (
	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FiledownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFiledownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FiledownloadLogic {
	return &FiledownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FiledownloadLogic) Filedownload(req *types.FileDownloadReq) (*types.FileDownloadResp, error) {

	//resp, err := l.svcCtx.FileRpc.FileUpload(l.ctx, &pb.FileUploadReq{
	//	Name:       req.Name,
	//	Type:       req.Type,
	//	SourcePath: req.SourcePath,
	//})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//if resp.Error != "" {
	//	return nil, errorx.NewDefaultError(resp.Error)
	//}

	return nil, nil
}

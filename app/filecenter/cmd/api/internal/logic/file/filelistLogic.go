package file

import (
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilelistLogic {
	return &FilelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilelistLogic) Filelist(req *types.FileListReq) (*types.FileListResp, error) {

	resp, err := l.svcCtx.FileRpc.FileList(l.ctx, &pb.FileListReq{
		Id:    req.Id,
		Page:  req.Page,
		Size:  req.Size,
		Field: req.Field,
		ASC:   req.ASC,
	})

	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errorx.NewDefaultError(resp.Error)
	}

	var res types.FileListResp

	for _, x := range resp.List {
		var y = types.File_in{
			Name:       x.Name,
			Type:       x.Type,
			Path:       x.Path,
			Size:       x.Size,
			ShareLink:  x.ShareLink,
			ModifyTime: x.UpdatedAt,
		}
		res.List = append(res.List, y)
	}
	res.Count = resp.Count

	return &res, nil
}

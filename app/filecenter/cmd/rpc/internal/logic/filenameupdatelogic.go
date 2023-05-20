package logic

import (
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileNameUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileNameUpdateLogic {
	return &FileNameUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileNameUpdateLogic) FileNameUpdate(in *pb.FileNameUpdateReq) (*pb.FileNameUpdateResp, error) {

	// 查询文件ID 更改文件名
	_, err := l.svcCtx.FileModel.UpdateOneMapById(l.ctx, in.Id, map[string]interface{}{
		"name": in.Name,
	})

	if err != nil {
		return nil, err
	}

	return &pb.FileNameUpdateResp{
		Error: "修改完成",
	}, nil
}

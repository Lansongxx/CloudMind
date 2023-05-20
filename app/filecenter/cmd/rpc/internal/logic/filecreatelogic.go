package logic

import (
	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/app/filecenter/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileCreateLogic {
	return &FileCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileCreateLogic) FileCreate(in *pb.FileCreateReq) (*pb.FileCreateResp, error) {

	_, err := l.svcCtx.FileModel.Insert(l.ctx, &model.File{
		//Id: 6,
		Name:      in.Name,
		Type:      "folder",
		Path:      "xxxx",
		Size:      "0",
		ShareLink: "xxxxxx",
		ParentId:  in.ParentId,
	})

	if err != nil {
		return &pb.FileCreateResp{
			Error: "创建失败",
		}, err
	}

	return &pb.FileCreateResp{
		Error: "创建成功",
	}, nil
}

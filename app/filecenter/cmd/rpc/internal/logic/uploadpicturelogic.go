package logic

import (
	"context"

	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPictureLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPictureLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPictureLogic {
	return &UploadPictureLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadPictureLogic) UploadPicture(in *pb.UploadPictureReq) (*pb.UploadPictureResp, error) {

	// 开启事务

	return &pb.UploadPictureResp{}, nil
}

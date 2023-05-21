package logic

import (
	"CloudMind/app/filecenter/model"
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

	/*
		in.Picture // 图片的字节数组
	*/

	resp, err := l.svcCtx.UserProfileModel.TxInsert(l.ctx, l.svcCtx.GormDB, &model.UserProfile{
		UserId: in.Id,
		Md5:    in.Md5,
	}, in.Picture)

	if err != nil {
		return nil, err
	}

	if resp == "" {
		return &pb.UploadPictureResp{
			Error: "上传失败",
		}, nil
	}

	return &pb.UploadPictureResp{
		URL: resp,
	}, nil
}

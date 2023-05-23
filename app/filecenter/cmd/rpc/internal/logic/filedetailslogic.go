package logic

import (
	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/app/filecenter/model"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDetailsLogic {
	return &FileDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDetailsLogic) FileDetails(in *pb.FileDetailsReq) (*pb.FileDetailsResp, error) {

	File, err := l.svcCtx.FileModel.FindOne(l.ctx, in.Id) // 查看文件详情

	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

	if err == model.ErrNotFound {
		return &pb.FileDetailsResp{
			Error: "文件不存在",
		}, nil
	}

	var pbFile pb.FileDetailsResp
	if File != nil {
		err = copier.Copy(&pbFile, File)

		if err != nil {
			return nil, err
		}
	}

	return &pbFile, nil
}

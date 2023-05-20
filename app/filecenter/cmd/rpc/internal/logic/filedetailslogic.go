package logic

import (
	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

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

	Filex, err := l.svcCtx.FileModel.FindOne(l.ctx, in.Id) // 查看文件详情

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("出错了")
	}

	var pbFile pb.FileDetailsResp

	if Filex != nil {
		err = copier.Copy(&pbFile, Filex)

		if err != nil {
			return nil, err
		}
	}
	return &pbFile, nil
}

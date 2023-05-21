package file

import (
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/common/errorx"
	"context"

	"CloudMind/app/filecenter/cmd/api/internal/svc"
	"CloudMind/app/filecenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FiledeletionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFiledeletionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FiledeletionLogic {
	return &FiledeletionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FiledeletionLogic) Filedeletion(req *types.FileDeletionReq) (*types.FileDeletionResp, error) {

	//var list []*pb.FilePrefix
	//for _, file := range req.Delist {
	//	list = append(list, &pb.FilePrefix{
	//		ParentId: file.ParentId,
	//		Id:       file.Id,
	//		Type:     file.Type,
	//	})
	//}

	resp, err := l.svcCtx.FileRpc.FileDeletion(l.ctx, &pb.FileDeletionReq{
		List: req.Delist,
	})

	if err != nil {
		return nil, err
	}

	if resp.Error != "" {
		return nil, errorx.NewDefaultError(resp.Error)
	}

	return nil, nil
}

package logic

import (
	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"CloudMind/app/filecenter/model"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileListLogic) FileList(in *pb.FileListReq) (*pb.FileListResp, error) {

	/*
			分页查询

			查询条件集合
			Query{
				Field     string  查询字段名
		    	Condition ConditionType 判断符号
		    	Value     interface{} 条件值
			}

			Page{
				PageIndex: int(in.Page), 页码
				PageSize:  int(in.Size), 该页最大展示记录数
			}

			Order{
				Field string  排序字段
		    	ASC   bool	  true为升序
			}

	*/

	resp, x, err := l.svcCtx.FileModel.FindsByPage(l.ctx, []model.Query{
		{"parent_id", "=", in.Id},
	}, &model.Page{
		PageIndex: int(in.Page),
		PageSize:  int(in.Size),
	}, []model.Order{
		{Field: in.Field, ASC: in.ASC},
	})

	if err != nil {
		return &pb.FileListResp{
			Error: "查询失败",
		}, err
	}

	var res pb.FileListResp
	res.Count = x // 实际记录条数

	for _, file := range resp {
		res.List = append(res.List, &pb.FileIn{
			Name:      file.Name,             // 文件名
			Type:      file.Type,             // 文件类型
			Path:      file.Path,             // 文件路径
			Size:      file.Size,             // 文件大小
			ShareLink: file.ShareLink,        // 分享链接
			UpdatedAt: file.CreatedAt.Unix(), // 最新更改时间戳
		})
	}

	return &res, nil
}

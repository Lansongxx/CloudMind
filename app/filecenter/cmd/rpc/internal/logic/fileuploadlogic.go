package logic

import (
	"CloudMind/app/filecenter/cmd/rpc/internal/svc"
	"CloudMind/app/filecenter/cmd/rpc/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadLogic) FileUpload(in *pb.FileUploadReq) (*pb.FileUploadResp, error) {

	//client, err := oss.New(l.svcCtx.Config.OSS.Conf.Endpoint, l.svcCtx.Config.OSS.Conf.AccessKeyID, l.svcCtx.Config.OSS.Conf.AccessKeySecret)
	//if err != nil {
	//	return &pb.FileUploadResp{
	//		Error: "上传失败",
	//	}, err
	//}
	//
	//// 填写Bucket名称，例如examplebucket。
	//bucket, err := client.Bucket(l.svcCtx.Config.OSS.Bucket)
	//if err != nil {
	//	return &pb.FileUploadResp{
	//		Error: "上传失败",
	//	}, err
	//}
	//
	//// 设置分片大小为100 KB（100*1024），指定分片上传并发数为3，并开启断点续传上传。
	//// yourObjectName填写Object完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	//// yourLocalFile填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。如果未指定本地路径，则默认从示例程序所属项目对应本地路径中上传文件。
	//err = bucket.UploadFile("exampledir/exampleobject.txt", "D:\\localpath\\examplefile.txt", 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}

	return &pb.FileUploadResp{}, nil
}

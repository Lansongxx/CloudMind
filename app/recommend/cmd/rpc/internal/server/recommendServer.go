// Code generated by goctl. DO NOT EDIT.
// Source: recommend.proto

package server

import (
	"context"

	"CloudMind/app/recommend/cmd/rpc/internal/logic"
	"CloudMind/app/recommend/cmd/rpc/internal/svc"
	"CloudMind/app/recommend/cmd/rpc/pb"
)

type RecommendServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRecommendServer
}

func NewRecommendServer(svcCtx *svc.ServiceContext) *RecommendServer {
	return &RecommendServer{
		svcCtx: svcCtx,
	}
}

func (s *RecommendServer) InsertFeedBack(ctx context.Context, in *pb.InsertFeedBackReq) (*pb.InsertFeedBackResp, error) {
	l := logic.NewInsertFeedBackLogic(ctx, s.svcCtx)
	return l.InsertFeedBack(in)
}

func (s *RecommendServer) InsertItem(ctx context.Context, in *pb.InsertItemReq) (*pb.InsertItemResp, error) {
	l := logic.NewInsertItemLogic(ctx, s.svcCtx)
	return l.InsertItem(in)
}

func (s *RecommendServer) DeleteItem(ctx context.Context, in *pb.DeleteItemReq) (*pb.DeleteItemResp, error) {
	l := logic.NewDeleteItemLogic(ctx, s.svcCtx)
	return l.DeleteItem(in)
}

func (s *RecommendServer) UpdateItem(ctx context.Context, in *pb.UpdateItemReq) (*pb.UpdateItemResp, error) {
	l := logic.NewUpdateItemLogic(ctx, s.svcCtx)
	return l.UpdateItem(in)
}

func (s *RecommendServer) GetRecommendByItem(ctx context.Context, in *pb.GetRecommendByItemIdReq) (*pb.GetRecommendByItemIdResp, error) {
	l := logic.NewGetRecommendByItemLogic(ctx, s.svcCtx)
	return l.GetRecommendByItem(in)
}

func (s *RecommendServer) GetRecommendByUserId(ctx context.Context, in *pb.GetRecommendByUserIdReq) (*pb.GetRecommendByUserIdResp, error) {
	l := logic.NewGetRecommendByUserIdLogic(ctx, s.svcCtx)
	return l.GetRecommendByUserId(in)
}

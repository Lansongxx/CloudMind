package logic

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/svc"
	"CloudMind/app/usercenter/cmd/rpc/pb"
	"CloudMind/app/usercenter/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if err != nil {
		generateToken := NewGenerateTokenLogic(l.ctx, l.svcCtx)
		_, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
			Email:      in.Email,
			Nickname:   in.NickName,
			Password:   in.PassWord,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			DeleteTime: time.Now(),
		})
		if err != nil {
			return nil, err
		}

		count, err := l.svcCtx.UserModel.FindCount(l.ctx, []model.Query{})

		TokenResp, err := generateToken.GenerateToken(&pb.GenerateTokenReq{
			UserId: count,
		})

		if err != nil {
			return nil, err
		}
		return &pb.RegisterResp{
			AccessToken:  TokenResp.AccessToken,
			AccessExpire: TokenResp.AccessExpire,
			RefreshAfter: TokenResp.RefreshAfter,
		}, nil

	}
	return nil, errors.New("该邮箱已经注册过，请勿重复注册")
}

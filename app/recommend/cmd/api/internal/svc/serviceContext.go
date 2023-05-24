package svc

import (
	"CloudMind/app/recommend/cmd/api/internal/config"
	"CloudMind/app/recommend/cmd/rpc/recommend"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	RecommendRpc recommend.Recommend
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		RecommendRpc: recommend.NewRecommend(zrpc.MustNewClient(c.RecommendRpcConf)),
	}
}

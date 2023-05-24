package svc

import (
	"CloudMind/app/recommend/cmd/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zhenghaoz/gorse/client"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	Gorse  *client.GorseClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Redis:  redis.MustNewRedis(c.RedisConf),
		Gorse:  client.NewGorseClient("http://master:8088", ""),
	}
}

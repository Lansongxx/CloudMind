package config

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	DB struct {
		DataSource string
	}

	Cache cache.CacheConf

	OSS struct {
		Conf   oss.Config
		Bucket string
	}
}

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	DB struct {
		DataSource string
	}

	Cache cache.CacheConf

	//OSS struct {
	//	Conf   oss.Config
	//	Bucket string
	//}
}

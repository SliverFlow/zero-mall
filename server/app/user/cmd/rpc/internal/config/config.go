package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"server/common/xconfig"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql      xconfig.Mysql
	CacheRedis cache.CacheConf

	ProductRpc zrpc.RpcClientConf
}

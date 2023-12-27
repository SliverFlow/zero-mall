package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"server/common/xconfig"
)

type Config struct {
	zrpc.RpcServerConf

	UserRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf

	Mysql xconfig.Mysql
}

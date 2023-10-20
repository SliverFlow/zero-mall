package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	CartRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf

	// jwt 相关配置
	XJwt struct {
		Isuser          string // 签发者
		Secret          string // 加密字符串
		Expire          int64  // 过期时间
		Buffer          int64  // 续期时间
		BlackListPrefix string // 黑名单前缀
	}

	Redis redis.RedisConf
}

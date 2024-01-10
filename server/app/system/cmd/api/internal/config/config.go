package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"server/common/xupload"
)

type Config struct {
	rest.RestConf

	// jwt 相关配置
	XJwt struct {
		Isuser          string // 签发者
		Secret          string // 加密字符串
		Expire          int64  // 过期时间
		Buffer          int64  // 续期时间
		BlackListPrefix string // 黑名单前缀
	}
	Redis redis.RedisConf

	StsAliOss xupload.StsConfig //  阿里云oss临时上传凭证配置

	Captcha struct {
		KeyLong   int
		ImgHeight int
		ImgWidth  int
	}
	SystemRpc  zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
	LogRpc     zrpc.RpcClientConf
}

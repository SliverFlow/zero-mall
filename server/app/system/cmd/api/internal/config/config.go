package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
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

	Oss struct {
		Type  string
		Local struct {
			Path string
		}
		Alioss struct {
			Endpoint        string
			AccessKeyId     string
			AccessKeySecret string
			BucketName      string
		}
	}
}

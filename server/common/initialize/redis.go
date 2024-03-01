package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"server/common/xconfig"
)

func InitRedis(c xconfig.Redis) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Addr(),
		Password: c.Password, // 没有密码，默认值
		DB:       c.Db,       // 默认DB 0
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	logx.Info("[REDIS] : ", "redis connect success")
	return rdb, nil
}

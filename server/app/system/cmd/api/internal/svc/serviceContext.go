package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"server/app/system/cmd/api/internal/config"
	"server/app/system/cmd/rpc/pb"
	"server/app/system/cmd/rpc/system"
	"server/common/middleware"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	SystemRpc system.System
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwtConf := c.XJwt
	client := NewRedisClient(c.Redis)

	return &ServiceContext{
		Config:    c,
		Auth:      middleware.NewAuthMiddleware(jwtConf.Secret, jwtConf.Expire, jwtConf.Buffer, jwtConf.Isuser, jwtConf.BlackListPrefix, client).Handle, // 认证中间件
		SystemRpc: NewSystemRpc(c.SystemRpc),
	}
}

func NewRedisClient(c redis2.RedisConf) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:         c.Host,
		Password:     c.Pass,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})
	result, err := cli.Ping(context.Background()).Result()
	if err != nil {
		os.Exit(0)
		return nil
	}
	logx.Info("[REDIS CONNECTION SUCCESS ] PING RESULT : ", result)
	return cli
}

func NewSystemRpc(c zrpc.RpcClientConf) pb.SystemClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] system rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] system rpc connection success : %v\n")
	return system.NewSystem(client)
}

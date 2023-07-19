package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"server/app/user/cmd/api/internal/config"
	"server/app/user/cmd/rpc/pb"
	"server/app/user/cmd/rpc/user"
	"server/common/middleware"
	"time"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	UserRpc user.User
	Rsc     *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	jwtConf := c.XJwt
	client := NewRedisClient(c.Redis)
	return &ServiceContext{
		Config:  c,
		Rsc:     client,
		Auth:    middleware.NewAuthMiddleware(jwtConf.Secret, jwtConf.Expire, jwtConf.Buffer, jwtConf.Isuser, jwtConf.BlackListPrefix, client).Handle, // 认证中间件
		UserRpc: NewUserRpc(c.UserRpc),
	}
}

// NewUserRpc
// Author [SliverFlow]
// @desc user rpc 连接
func NewUserRpc(c zrpc.RpcClientConf) pb.UserClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] user rpc client conn err: %v\n ", err)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS	] user rpc client conn success \n")
	return user.NewUser(client)
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
	logx.Info("[REDIS CONNECTION SUCCESS ] PING RESULT : %v\n", result)
	return cli
}

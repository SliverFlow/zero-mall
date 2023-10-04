package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"server/app/order/cmd/api/internal/config"
	"server/app/order/cmd/rpc/order"
	"server/common/middleware"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {

	jwtConf := c.XJwt
	client := NewRedisClient(c.Redis)

	return &ServiceContext{
		Config:   c,
		Auth:     middleware.NewAuthMiddleware(jwtConf.Secret, jwtConf.Expire, jwtConf.Buffer, jwtConf.Isuser, jwtConf.BlackListPrefix, client).Handle,
		OrderRpc: newOrderRpc(c.OrderRpc),
	}
}

func newOrderRpc(c zrpc.RpcClientConf) order.Order {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] order rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] order rpc connection success : %v\n")
	return order.NewOrder(client)
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

package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"server/app/order/cmd/rpc/order"
	orderpb "server/app/order/cmd/rpc/pb"
	pruductpb "server/app/product/cmd/rpc/pb"
	"server/app/product/cmd/rpc/product"
	"server/app/system/cmd/api/internal/config"
	systempb "server/app/system/cmd/rpc/pb"
	"server/app/system/cmd/rpc/system"
	userpb "server/app/user/cmd/rpc/pb"
	"server/app/user/cmd/rpc/user"
	"server/common/middleware"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	SystemRpc  system.System
	UserRpc    user.User
	ProductRpc product.Product
	OrderRpc   order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwtConf := c.XJwt
	client := NewRedisClient(c.Redis)

	return &ServiceContext{
		Config:     c,
		Auth:       middleware.NewAuthMiddleware(jwtConf.Secret, jwtConf.Expire, jwtConf.Buffer, jwtConf.Isuser, jwtConf.BlackListPrefix, client).Handle, // 认证中间件
		SystemRpc:  newSystemRpc(c.SystemRpc),
		UserRpc:    newUserRpc(c.UserRpc),
		ProductRpc: newProductRpc(c.ProductRpc),
		OrderRpc:   newOrderRpc(c.OrderRpc),
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

// 强依赖 system rpc
func newSystemRpc(c zrpc.RpcClientConf) systempb.SystemClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] system rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] system rpc connection success : %v\n")
	return system.NewSystem(client)
}

// 强依赖 user rpc
func newUserRpc(c zrpc.RpcClientConf) userpb.UserClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] user rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] user rpc connection success : %v\n")
	return user.NewUser(client)
}

// 强依赖 product rpc
func newProductRpc(c zrpc.RpcClientConf) pruductpb.ProductClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] product rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] product rpc connection success : %v\n")
	return product.NewProduct(client)
}

// 强依赖 order rpc
func newOrderRpc(c zrpc.RpcClientConf) orderpb.OrderClient {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] order rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] order rpc connection success : %v\n")
	return order.NewOrder(client)
}

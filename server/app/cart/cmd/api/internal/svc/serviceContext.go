package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"server/app/cart/cmd/api/internal/config"
	"server/app/cart/cmd/rpc/cart"
	"server/app/product/cmd/rpc/product"
	"server/common/middleware"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	CartRpc    cart.Cart
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {

	jwtConf := c.XJwt
	client := NewRedisClient(c.Redis)

	return &ServiceContext{
		Config:     c,
		Auth:       middleware.NewAuthMiddleware(jwtConf.Secret, jwtConf.Expire, jwtConf.Buffer, jwtConf.Isuser, jwtConf.BlackListPrefix, client).Handle,
		CartRpc:    newCartRpc(c.CartRpc),
		ProductRpc: newProductRpc(c.ProductRpc),
	}
}

// 强依赖 cart rpc
func newCartRpc(c zrpc.RpcClientConf) cart.Cart {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] cart rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] cart rpc connection success : %v\n")
	return cart.NewCart(client)
}

// 强依赖 cart rpc
func newProductRpc(c zrpc.RpcClientConf) product.Product {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] product rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] product rpc connection success : %v\n")
	return product.NewProduct(client)
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

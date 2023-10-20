package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"server/app/order/cmd/api/internal/config"
	"server/app/order/cmd/rpc/order"
	"server/app/product/cmd/rpc/product"
	"server/app/user/cmd/rpc/user"
	"server/common/middleware"
	"time"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	OrderRpc   order.Order
	ProductRpc product.Product
	UserRpc    user.User
	EtcdCli    *clientv3.Client
}

func NewServiceContext(c config.Config) *ServiceContext {

	jwtConf := c.XJwt
	client := NewRedisClient(c.Redis)

	return &ServiceContext{
		Config:     c,
		Auth:       middleware.NewAuthMiddleware(jwtConf.Secret, jwtConf.Expire, jwtConf.Buffer, jwtConf.Isuser, jwtConf.BlackListPrefix, client).Handle,
		OrderRpc:   newOrderRpc(c.OrderRpc),
		ProductRpc: newProductRpc(c.ProductRpc),
		UserRpc:    newUserRpc(c.UserRpc),
		EtcdCli:    newEtcdClient(c.EtcdLocker.Hosts),
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

func newUserRpc(c zrpc.RpcClientConf) user.User {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] user rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS ] user rpc connection success : %v\n")
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
	logx.Info("[REDIS CONNECTION SUCCESS ] PING RESULT : ", result)
	return cli
}

func newEtcdClient(endpoints []string) *clientv3.Client {
	c, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	if err != nil {
		logx.Errorf("[ETCD CONNECTION ERROR] ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[ETCD CONNECTION SUCCESS]")
	return c
}

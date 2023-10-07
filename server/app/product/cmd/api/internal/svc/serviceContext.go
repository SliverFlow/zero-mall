package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"server/app/product/cmd/api/internal/config"
	"server/app/product/cmd/api/internal/middleware"
	"server/app/product/cmd/rpc/product"
	"server/app/user/cmd/rpc/user"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware

	ProductRpc product.Product
	UserRpc    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Auth:       middleware.NewAuthMiddleware().Handle,
		ProductRpc: newProductRpc(c.ProductRpc),
		UserRpc:    newUserRpc(c.UserRpc),
	}
}

func newProductRpc(c zrpc.RpcClientConf) product.Product {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] product rpc client conn err: %v\n ", err)
		os.Exit(0)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS	] product rpc client conn success \n")
	return product.NewProduct(client)
}

func newUserRpc(c zrpc.RpcClientConf) user.User {
	client, err := zrpc.NewClient(c)
	if err != nil {
		logx.Errorf("[RPC CONNECTION ERROR] user rpc client conn err: %v\n ", err)
		return nil
	}
	logx.Info("[RPC CONNECTION SUCCESS	] user rpc client conn success \n")
	return user.NewUser(client)
}

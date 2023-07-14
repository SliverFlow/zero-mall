package svc

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"server/app/user/cmd/api/internal/config"
	"server/app/user/cmd/rpc/pb"
	"server/app/user/cmd/rpc/user"
	"server/common/middleware"
)

type ServiceContext struct {
	Config  config.Config
	Auth    rest.Middleware
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:  c,
		Auth:    middleware.NewAuthMiddleware().Handle, // 认证中间件
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

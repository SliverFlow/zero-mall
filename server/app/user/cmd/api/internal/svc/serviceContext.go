package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"server/app/user/cmd/api/internal/config"
	"server/common/middleware"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle, // 认证中间件
	}
}

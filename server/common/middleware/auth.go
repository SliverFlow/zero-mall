package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"server/common/result"
	"server/common/xerr"
)

// Auth
// Author [SliverFlow]
// @desc 全局的认证中间件
type Auth struct {
}

func NewAuthMiddleware() *Auth {
	return &Auth{}
}

func (m *Auth) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("request for auth middleware")
		authorization := r.Header.Get("Authorization")
		if authorization == "" { // 为获取到 token
			result.HttpResult(r, w, nil, xerr.NewCodeError(200001))
			return
		}
		// 解析 token

		// Passthrough to next handler if need
		next(w, r)
	}
}

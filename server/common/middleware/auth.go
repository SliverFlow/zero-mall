package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"server/common/result"
	"server/common/xerr"
	"server/common/xjwt"
)

// Auth
// Author [SliverFlow]
// @desc 全局的认证中间件
type Auth struct {
	Isuser string // 签发者
	Secret string // 加密字符串
	Expire int64  // 过期时间
	Buffer int64  // 续期时间
}

func NewAuthMiddleware(secret string, expire, buffer int64, isuser string) *Auth {
	return &Auth{
		Isuser: isuser,
		Secret: secret,
		Expire: expire,
		Buffer: buffer,
	}
}

func (m *Auth) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("request for auth middleware")
		authorization := r.Header.Get("Authorization")
		if authorization == "" { // 未获取到 token
			result.HttpResult(r, w, nil, xerr.NewErrMsg("非法请求，请登录后再使用"))
			return
		}
		// 解析 token
		jwt := xjwt.NewXJwt([]byte(m.Secret), m.Expire, m.Buffer, m.Isuser)
		_, err := jwt.ParseToken(authorization)
		if err != nil {
			result.HttpResult(r, w, nil, xerr.NewErrMsg(err.Error()))
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}

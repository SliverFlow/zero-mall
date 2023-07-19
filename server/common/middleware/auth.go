package middleware

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"server/common/result"
	"server/common/xerr"
	"server/common/xjwt"
	"strconv"
	"time"
)

// Auth
// Author [SliverFlow]
// @desc 全局的认证中间件
type Auth struct {
	isuser          string // 签发者
	secret          string // 加密字符串
	expire          int64  // 过期时间
	buffer          int64  // 续期时间
	blackListPrefix string
	rsc             *redis.Client
}

func NewAuthMiddleware(secret string, expire, buffer int64, isuser, blackListPrefix string, rsc *redis.Client) *Auth {
	return &Auth{
		isuser:          isuser,
		secret:          secret,
		expire:          expire,
		buffer:          buffer,
		blackListPrefix: blackListPrefix,
		rsc:             rsc,
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
		jwt := xjwt.NewXJwt([]byte(m.secret), m.expire, m.buffer, m.isuser, m.blackListPrefix)
		// 判断 token 是不是在黑名单中
		ok := jwt.IsRedisBlackList(authorization, m.rsc)
		if ok {
			result.HttpResult(r, w, nil, xerr.NewErrMsg("黑名单或异地登录"))
			return
		}

		claims, err := jwt.ParseToken(authorization)
		if err != nil {
			result.HttpResult(r, w, nil, xerr.NewErrMsg(err.Error()))
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime { // 更换 token
			token, _ := jwt.Renewal(authorization)
			// 新 token
			w.Header().Set("new-token", token)
			// 新的 token 过期时间
			w.Header().Set("new-expire", strconv.FormatInt(time.Now().Unix()+claims.BufferTime, 10))
			// 将旧的 token 添加进黑名单
			_ = jwt.RedisBlackList(authorization, m.rsc)
		}
		// 传递 claims
		ctx := context.WithValue(r.Context(), "claims", claims)
		// Passthrough to next handler if need
		next(w, r.WithContext(ctx))
	}
}

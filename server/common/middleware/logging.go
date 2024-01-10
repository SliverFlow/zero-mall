package middleware

import (
	"net"
	"net/http"
	logpb "server/app/log/cmd/rpc/pb"
	userpb "server/app/user/cmd/rpc/pb"
	"server/common/xjwt"
	"strings"
	"time"
)

type Logging struct {
	userRpc userpb.UserClient
	logRpc  logpb.LogClient
}

func NewLoggingMiddleware(userRpc userpb.UserClient, logRpc logpb.LogClient) *Logging {
	return &Logging{
		userRpc: userRpc,
		logRpc:  logRpc,
	}
}

func (l *Logging) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uuid, err := xjwt.GetUserUUID(r.Context())
		if err != nil {
			return // 无需记录日志
		}

		userpbrep, err := l.userRpc.UserFindByUUID(r.Context(), &userpb.UUIDReq{UUID: uuid})

		ctx := r.Context()
		timer := ctx.Value("timer")

		next(w, r.WithContext(r.Context()))
		unix := time.Now().UnixNano() / int64(time.Millisecond)
		_, _ = l.logRpc.LogCreate(r.Context(), &logpb.LogCreateReq{
			UserID:        uuid,
			Username:      userpbrep.User.Nickname + "(" + userpbrep.User.Username + ")",
			IP:            getClientIP(r),
			Method:        r.Method,
			Path:          r.URL.Path,
			RequestParams: r.URL.RawQuery,
			Latency:       unix - timer.(int64),
		})
	}
}

func getClientIP(r *http.Request) string {
	// 尝试从 X-Real-IP 头字段获取 IP 地址
	ip := r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}
	// 尝试从 X-Forwarded-For 头字段获取 IP 地址
	ip = r.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For 可能包含多个 IP 地址，第一个是客户端的实际 IP 地址
		if index := strings.IndexByte(ip, ','); index >= 0 {
			ip = ip[0:index]
		}
		return ip
	}
	// 从 RemoteAddr 获取 IP 地址，去掉端口号
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return "unknown"
	}
	return ip
}

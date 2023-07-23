package base

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/system/cmd/api/internal/logic/base"
	"server/app/system/cmd/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload()
		result.HttpResult(r, w, resp, err)
	}
}

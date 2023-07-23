package base

import (
	"net/http"
	"server/common/result"

	"server/app/system/cmd/api/internal/logic/base"
	"server/app/system/cmd/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(r)
		result.HttpResult(r, w, resp, err)
	}
}

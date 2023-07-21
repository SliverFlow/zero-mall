package public

import (
	"net/http"
	"server/app/system/cmd/api/internal/logic/public"
	"server/common/result"

	"server/app/system/cmd/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(r)
		result.HttpResult(r, w, resp, err)
	}
}

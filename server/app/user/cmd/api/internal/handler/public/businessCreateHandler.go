package public

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/user/cmd/api/internal/logic/public"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
)

func BusinessCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBusinessReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := public.NewBusinessCreateLogic(r.Context(), svcCtx)
		resp, err := l.BusinessCreate(&req)
		result.HttpResult(r, w, resp, err)
	}
}

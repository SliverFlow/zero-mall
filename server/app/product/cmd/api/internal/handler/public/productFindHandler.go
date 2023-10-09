package public

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/product/cmd/api/internal/logic/public"
	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"
)

func ProductFindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductFindReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := public.NewProductFindLogic(r.Context(), svcCtx)
		resp, err := l.ProductFind(&req)
		result.HttpResult(r, w, resp, err)
	}
}

package public

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/product/cmd/api/internal/logic/public"
	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"
)

func CategoryIDByProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryIDByProductListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := public.NewCategoryIDByProductListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryIDByProductList(&req)
		result.HttpResult(r, w, resp, err)
	}
}

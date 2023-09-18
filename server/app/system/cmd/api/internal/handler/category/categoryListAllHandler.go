package category

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/system/cmd/api/internal/logic/category"
	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"
)

func CategoryListAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NilReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := category.NewCategoryListAllLogic(r.Context(), svcCtx)
		resp, err := l.CategoryListAll(&req)
		result.HttpResult(r, w, resp, err)
	}
}

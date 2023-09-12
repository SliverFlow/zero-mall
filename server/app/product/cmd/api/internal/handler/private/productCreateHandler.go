package private

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/product/cmd/api/internal/logic/private"
	"server/app/product/cmd/api/internal/svc"
	"server/app/product/cmd/api/internal/types"
)

func ProductCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := private.NewProductCreateLogic(r.Context(), svcCtx)
		resp, err := l.ProductCreate(&req)
		result.HttpResult(r, w, resp, err)
	}
}

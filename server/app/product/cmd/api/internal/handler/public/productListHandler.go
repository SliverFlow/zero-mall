package public

import (
	"net/http"
	"server/common/result"

	"server/app/product/cmd/api/internal/logic/public"
	"server/app/product/cmd/api/internal/svc"
)

func ProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewProductListLogic(r.Context(), svcCtx)
		err := l.ProductList()
		result.HttpResult(r, w, nil, err)
	}
}

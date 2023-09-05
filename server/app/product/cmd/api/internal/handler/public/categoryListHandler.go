package public

import (
	"net/http"
	"server/common/result"

	"server/app/product/cmd/api/internal/logic/public"
	"server/app/product/cmd/api/internal/svc"
)

func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewCategoryListLogic(r.Context(), svcCtx)
		err := l.CategoryList()
		result.HttpResult(r, w, nil, err)
	}
}

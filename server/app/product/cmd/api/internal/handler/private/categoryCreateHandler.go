package private

import (
	"net/http"
	"server/common/result"

	"server/app/product/cmd/api/internal/logic/private"
	"server/app/product/cmd/api/internal/svc"
)

func CategoryCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := private.NewCategoryCreateLogic(r.Context(), svcCtx)
		err := l.CategoryCreate()
		result.HttpResult(r, w, nil, err)
	}
}

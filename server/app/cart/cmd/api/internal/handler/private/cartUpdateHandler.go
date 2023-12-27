package private

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/cart/cmd/api/internal/logic/private"
	"server/app/cart/cmd/api/internal/svc"
	"server/app/cart/cmd/api/internal/types"
)

func CartUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := private.NewCartUpdateLogic(r.Context(), svcCtx)
		resp, err := l.CartUpdate(&req)
		result.HttpResult(r, w, resp, err)
	}
}

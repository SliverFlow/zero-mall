package private

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/order/cmd/api/internal/logic/private"
	"server/app/order/cmd/api/internal/svc"
	"server/app/order/cmd/api/internal/types"
)

func OrderFindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Nil
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := private.NewOrderFindLogic(r.Context(), svcCtx)
		resp, err := l.OrderFind(&req)
		result.HttpResult(r, w, resp, err)
	}
}

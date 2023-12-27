package private

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/order/cmd/api/internal/logic/private"
	"server/app/order/cmd/api/internal/svc"
	"server/app/order/cmd/api/internal/types"
)

func OrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := private.NewOrderListLogic(r.Context(), svcCtx)
		resp, err := l.OrderList(&req)
		result.HttpResult(r, w, resp, err)
	}
}

package business

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/system/cmd/api/internal/logic/business"
	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"
)

func BusinessListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := business.NewBusinessListLogic(r.Context(), svcCtx)
		resp, err := l.BusinessList(&req)
		result.HttpResult(r, w, resp, err)
	}
}

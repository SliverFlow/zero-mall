package role

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/system/cmd/api/internal/logic/role"
	"server/app/system/cmd/api/internal/svc"
	"server/app/system/cmd/api/internal/types"
)

func RoleCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := role.NewRoleCreateLogic(r.Context(), svcCtx)
		resp, err := l.RoleCreate(&req)
		result.HttpResult(r, w, resp, err)
	}
}

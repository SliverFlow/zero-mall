package private

import (
	"net/http"
	"server/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"server/app/user/cmd/api/internal/logic/private"
	"server/app/user/cmd/api/internal/svc"
	"server/app/user/cmd/api/internal/types"
)

func UserFindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := private.NewUserFindLogic(r.Context(), svcCtx)
		resp, err := l.UserFind(&req)
		result.HttpResult(r, w, resp, err)
	}
}

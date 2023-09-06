package result

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"server/common/xerr"
)

// HttpResult 统一使用这个 http 返回方法
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		// 错误返回
		errCode := xerr.SERVER_COMMON_ERROR
		errMessage := "服务器开小差了，请稍后重试"
		caseErr := errors.Cause(err)
		if e, ok := caseErr.(*xerr.CodeError); ok { // 将 err 断言
			errCode = e.GetErrCode()
			errMessage = e.GetErrMessage()
		} else {
			if gstatus, ok := status.FromError(caseErr); ok { // grpc err 错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { // 区分自定义错误以及底层其他的错误，例如 db 的错误等等
					errCode = grpcCode
					errMessage = gstatus.Message()
				}
			}
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		// 请求成功返回，给予前端业务错误
		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMessage))
	} else {
		// 成功返回
		r := Success(resp)
		// 将 r 写出前端
		httpx.WriteJson(w, http.StatusOK, r)
	}
}

// ParamErrorResult http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(xerr.REUQEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusOK, Error(xerr.REUQEST_PARAM_ERROR, errMsg))
}

package result

// ResponseSuccessBean 成功统一返回
type ResponseSuccessBean struct {
	Code    uint32      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// ResponseErrorBean   失败统一返回
type ResponseErrorBean struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

// Success 成功返回
func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{
		Code:    0,
		Message: "OK",
		Data:    data,
	}
}

// Error 失败返回
func Error(errCode uint32, errMessage string) *ResponseErrorBean {
	return &ResponseErrorBean{Code: errCode, Message: errMessage}
}

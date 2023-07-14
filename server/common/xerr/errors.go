package xerr

import "fmt"

// CodeError 自定义 error
type CodeError struct {
	errCode    uint32
	errMessage string
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMessage)
}

func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

func (e *CodeError) GetErrMessage() string {
	return e.errMessage
}

func NewCodeError(code uint32) *CodeError {
	return &CodeError{
		errCode:    code,
		errMessage: MapErrMsg(code),
	}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMessage: errMsg}
}

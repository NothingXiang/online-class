package resp

import (
	"fmt"
)

type APIError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (A *APIError) Error() string {

	return A.Msg
}

func newAPIError(code int, msg string) *APIError {
	return &APIError{Code: code, Msg: msg}
}

func (A *APIError) String() string {
	return A.Msg
}

// 在原来的基础上生成新的error
func (A *APIError) NewErr(err error) *APIError {
	return &APIError{
		Code: A.Code,
		Msg:  fmt.Sprintf("%v:%v", A.Msg, err),
	}
}

// 拼接错误信息，生成一个新的错误响应
func (A *APIError) NewErrStr(s string) *APIError {
	return &APIError{
		Code: A.Code,
		Msg:  fmt.Sprintf("%v:%v", A.Msg, s),
	}
}

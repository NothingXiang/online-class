package resp

import (
	"github.com/NothingXiang/online-class/common/utils"
)

type APIError struct {
	Code    int
	Key     string
	Message string
}

func (A *APIError) Error() string {
	if utils.IsEmptyString(A.Message) {
		return A.Key
	}
	return A.Key + ":" + A.Message
}

func newAPIError(code int, key string) *APIError {
	return &APIError{Code: code, Key: key}
}

func (A *APIError) String() string {
	if utils.IsEmptyString(A.Message) {
		return A.Key
	}
	return A.Key + ":" + A.Message
}

func (A *APIError) SetMsg(err error) *APIError {
	A.Message = err.Error()
	return A
}

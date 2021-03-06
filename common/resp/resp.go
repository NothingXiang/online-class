package resp

import (
	"fmt"
)

type APIResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 用于返回带有处理结果和错误信息的resp
func NewAPIResp(data interface{}, err *APIError) *APIResp {
	return &APIResp{Code: err.Code, Msg: err.Error(), Data: data}
}

// 用于返回只带有错误信息的resp
func ErrResp(err error) *APIResp {
	switch err := err.(type) {
	case *APIError:
		return &APIResp{Code: err.Code, Msg: err.Error()}
	default:
		return ErrResp(UnknownError.NewErr(err))
	}

}

// 返回带有正确处理结果的resp
func NewSucResp(data interface{}) *APIResp {
	return NewAPIResp(data, NoError)
}

func (a *APIResp) SetData(data interface{}) *APIResp {
	a.Data = data
	return a
}

func (a *APIResp) SetMessage(err error) *APIResp {
	a.Msg = fmt.Sprintf("%v:%v", a.Msg, err.Error())
	return a
}

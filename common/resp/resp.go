package resp

type APIResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 用于返回带有处理结果和错误信息的resp
func NewAPIResp(data interface{}, err *APIError) *APIResp {
	return &APIResp{Code: err.Code, Msg: err.Message, Data: data}
}

// 用于返回只带有错误信息的resp
func NewErrResp(err *APIError) *APIResp {
	return &APIResp{Code: err.Code, Msg: err.Message}
}

// 返回带有正确处理结果的resp
func NewSucResp(data interface{}) *APIResp {
	return NewAPIResp(data, Success)
}

func (a *APIResp) SetData(data interface{}) *APIResp {
	a.Data = data
	return a
}

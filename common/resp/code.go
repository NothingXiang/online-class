package resp

var (
	// 这个不是error,这个是正确返回结果
	NoError = newAPIError(20000, "")

	UnknownError    = newAPIError(10001, "unknown error")
	NotExistError   = newAPIError(10002, "not exist")
	InvalidParamErr = newAPIError(10003, "invalid param")
	ParamFmtErr     = newAPIError(10004, "param format error")
	ParamEmptyErr   = newAPIError(10005, "param empty error")
	RepeatError     = newAPIError(10006, "repeat error")

	NotAuthError    = newAPIError(40001, "User authentication failed")
	TooManyReqError = newAPIError(40029, "too many request, has been limit")

	// other Errors ...
)

var (
	// 正确返回的结果
	Suc = NewAPIResp(nil, NoError)

	Unknown      = NewAPIResp(nil, UnknownError)
	NotExist     = NewAPIResp(nil, NotExistError)
	InvalidParam = NewAPIResp(nil, InvalidParamErr)
	ParamFmt     = NewAPIResp(nil, ParamFmtErr)
	NotAuth      = NewAPIResp(nil, NotAuthError)
	TooManyReq   = NewAPIResp(nil, TooManyReqError)
)

package resp

var (
	// 这个不是error,这个是正确返回结果
	Success = newAPIError(20000, "")

	UnknownError      = newAPIError(10001, "unknown errors")
	NotExistError     = newAPIError(10002, "not exist")
	InvalidParamError = newAPIError(10003, "invalid param")
	ParamFmtError     = newAPIError(10004, "param format errors")
	NotAuthError      = newAPIError(40001, "User authentication failed")
	TooManyReqError   = newAPIError(40029, "too many request, has been limit")

	// other Errors ...
)

var (
	// 正确返回的结果
	Normal = NewAPIResp(nil, Success)
	/*
		Unknown      = NewAPIResp(nil, Success)
		NotExist     = NewAPIResp(nil, NotExistError)
		InvalidParam = NewAPIResp(nil, InvalidParamError)
		ParamFmt     = NewAPIResp(nil, ParamFmtError)
		NotAuth      = NewAPIResp(nil, NotAuthError)
		TooManyReq   = NewAPIResp(nil, TooManyReqError)*/
)

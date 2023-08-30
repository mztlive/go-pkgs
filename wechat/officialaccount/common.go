package officialaccount

// CommonError 微信返回的通用错误json
type CommonError struct {
	apiName string
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

package officialaccount

// Config 返回给用户jssdk配置信息
type JsConfig struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

type Js interface {
	GetJsConfig(uri string) (config *JsConfig, err error)
}

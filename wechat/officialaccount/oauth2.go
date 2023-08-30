package officialaccount

// WebOauthAccessToken 获取用户授权access_token的返回结果
type WebOauthAccessToken struct {
	CommonError

	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`

	// IsSnapShotUser 是否为快照页模式虚拟账号，只有当用户是快照页模式虚拟账号时返回，值为1
	// 公众号文档 https://developers.weixin.qq.com/community/minihome/doc/000c2c34068880629ced91a2f56001
	IsSnapShotUser int `json:"is_snapshotuser"`

	// UnionID 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	// 公众号文档 https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842
	UnionID string `json:"unionid"`
}

type IOfficialAccountOauth interface {
	// GetRedirectURL 获取跳转的url地址
	GetRedirectURL(redirectURI, scope, state string) (string, error)

	// GetUserAccessToken 获取用户授权access_token
	GetUserAccessToken(code string) (result *WebOauthAccessToken, err error)
}

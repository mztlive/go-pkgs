package officialaccount

type OfficialAccountAccessToken interface {

	// GetAccessToken 获取公众号access_token
	GetAccessToken() (string, error)
}

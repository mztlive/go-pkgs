package wrapper

import (
	"github.com/jinzhu/copier"
	"github.com/mztlive/project-template/pkg/wechat/officialaccount"
)

// 这个文件实现了公众号授权相关的接口

// GetRedirectURL 获取授权链接
func (m *SilenceperOfficialAccount) GetRedirectURL(redirectURI, scope, state string) (string, error) {
	return m.engine.GetOauth().GetRedirectURL(redirectURI, scope, state)
}

// GetUserAccessToken 获取用户的Token
func (m *SilenceperOfficialAccount) GetUserAccessToken(code string) (result *officialaccount.WebOauthAccessToken, err error) {

	result = &officialaccount.WebOauthAccessToken{}

	accessToken, err := m.engine.GetOauth().GetUserAccessToken(code)
	if err != nil {
		return nil, err
	}

	copier.Copy(result, accessToken)
	return result, nil
}

package wrapper

import (
	"github.com/jinzhu/copier"
	"github.com/mztlive/go-pkgs/wechat/officialaccount"
)

func (m *SilenceperOfficialAccount) GetJsConfig(uri string) (config *officialaccount.JsConfig, err error) {
	config = &officialaccount.JsConfig{}
	cfg, err := m.engine.GetJs().GetConfig(uri)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(config, cfg)
	return
}

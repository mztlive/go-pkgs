package wrapper

import (
	"github.com/mztlive/go-pkgs/wechat/officialaccount"
	silenceper_official "github.com/silenceper/wechat/v2/officialaccount"
)

// SilenceperOfficialAccount 公众号订阅消息管理器
// 实现了IMiniSubscribeMessageManager接口
type SilenceperOfficialAccount struct {
	engine *silenceper_official.OfficialAccount

	messageHandler officialaccount.MessageHandlerFunc
}

func NewSilenceperOfficialAccountWrap(engine *silenceper_official.OfficialAccount) *SilenceperOfficialAccount {
	return &SilenceperOfficialAccount{
		engine: engine,
	}
}

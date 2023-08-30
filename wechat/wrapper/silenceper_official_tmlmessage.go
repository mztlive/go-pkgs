package wrapper

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mztlive/go-pkgs/wechat/officialaccount"
	silenceper_official "github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
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

// Send 发送模板消息
func (m *SilenceperOfficialAccount) Send(ctx context.Context, msg officialaccount.TemplateMessage) (int64, error) {

	message := message.TemplateMessage{}

	err := copier.Copy(&message, &msg)
	if err != nil {
		return 0, err
	}

	return m.engine.GetTemplate().Send(&message)
}

// AvailableTemplates	获取公众号账号下的模板列表
func (m *SilenceperOfficialAccount) AvailableTemplates(ctx context.Context) ([]officialaccount.TemplateItem, error) {
	list, err := m.engine.GetTemplate().List()
	if err != nil {
		return nil, fmt.Errorf("can not get template list: %w", err)
	}

	templates := make([]officialaccount.TemplateItem, 0, len(list))
	for _, item := range list {
		templates = append(templates, officialaccount.TemplateItem{
			TemplateID:      item.TemplateID,
			Title:           item.Title,
			PrimaryIndustry: item.PrimaryIndustry,
			DeputyIndustry:  item.DeputyIndustry,
			Content:         item.Content,
			Example:         item.Example,
		})
	}

	return templates, err
}

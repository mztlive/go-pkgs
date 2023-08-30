package wrapper

import (
	"context"

	"github.com/mztlive/project-template/pkg/wechat/mini"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	"github.com/spf13/cast"
)

// SilenceperMini 小程序订阅消息管理器
// 实现了IMiniSubscribeMessageManager接口
type SilenceperMini struct {
	engine *miniprogram.MiniProgram
}

func NewSilenceperMini(engine *miniprogram.MiniProgram) *SilenceperMini {
	return &SilenceperMini{
		engine: engine,
	}
}

// Send 发送订阅消息
func (m *SilenceperMini) Send(ctx context.Context, openID string, message mini.SubscribeMessage) error {
	mini := m.engine

	datas := make(map[string]*subscribe.DataItem, len(message.Data))
	for k, v := range message.Data {
		datas[k] = &subscribe.DataItem{Value: v}
	}

	return mini.GetSubscribe().Send(
		&subscribe.Message{
			ToUser:           openID,
			TemplateID:       message.TemplateID,
			Data:             datas,
			Page:             message.Page,
			MiniprogramState: message.MiniprogramState,
			Lang:             message.Lang,
		},
	)
}

// AvailableTemplates 获取可用的订阅消息模板
func (m *SilenceperMini) AvailableTemplates(ctx context.Context) ([]mini.SubscribeMessageTemplate, error) {

	list, err := m.engine.GetSubscribe().ListTemplates()
	if err != nil {
		return nil, err
	}

	templates := make([]mini.SubscribeMessageTemplate, 0, len(list.Data))
	for _, item := range list.Data {
		template := mini.SubscribeMessageTemplate{
			PriTmplId: item.PriTmplID,
			Content:   item.Content,
			Title:     item.Title,
			Example:   item.Example,
			Type:      cast.ToInt(item.Type),
		}

		templates = append(templates, template)
	}

	return templates, nil
}

package mini

import (
	"context"
)

// SubscribeMessage 订阅消息
type SubscribeMessage struct {
	TemplateID       string         `json:"template_id"`       // 必选，所需下发的订阅模板id
	Page             string         `json:"page"`              // 可选，点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             map[string]any `json:"data"`              // 必选, 模板内容
	MiniprogramState string         `json:"miniprogram_state"` // 可选，跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string         `json:"lang"`              // 入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

// SubscribeMessageTemplate 订阅消息模板
type SubscribeMessageTemplate struct {
	PriTmplId string `json:"priTmplId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	Type      int    `json:"type"`
}

// IMiniSubscribeMessageManager 订阅消息接口
type IMiniSubscribeMessageManager interface {

	// Send 发送订阅消息
	Send(ctx context.Context, openID string, message SubscribeMessage) error

	// AvailableTemplates	获取小程序账号下的模板列表
	AvailableTemplates(ctx context.Context) ([]SubscribeMessageTemplate, error)
}

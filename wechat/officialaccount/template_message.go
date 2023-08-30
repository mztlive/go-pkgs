package officialaccount

import "context"

type TemplateMessage struct {
	ToUser      string                       `json:"touser"`                  // 必须, 接受者OpenID
	TemplateID  string                       `json:"template_id"`             // 必须, 模版ID
	URL         string                       `json:"url,omitempty"`           // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color       string                       `json:"color,omitempty"`         // 可选, 整个消息的颜色, 可以不设置
	Data        map[string]*TemplateDataItem `json:"data"`                    // 必须, 模板数据
	ClientMsgID string                       `json:"client_msg_id,omitempty"` // 可选, 防重入ID

	MiniProgram struct {
		AppID    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` // 可选,跳转至小程序地址
}

// TemplateDataItem 模版内某个 .DATA 的值
type TemplateDataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// TemplateItem 模板消息.
type TemplateItem struct {
	TemplateID      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

// IOfficiAccountTemplateMessageManager 模板消息接口
type IOfficiAccountTemplateMessageManager interface {

	// Send 发送订阅消息
	Send(ctx context.Context, message TemplateMessage) (int64, error)

	// AvailableTemplates	获取公众号账号下的模板列表
	AvailableTemplates(ctx context.Context) ([]TemplateItem, error)
}

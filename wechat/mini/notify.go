package mini

const (
	EventTypeSubscribeMsgPopup = "subscribe_msg_popup_event" // 订阅消息弹框事件
)

const (
	// 消息类型是时间
	MessageTypeIsEvent = "event"
)

const (
	SubscribeStatusIsAccept = "accept" // 接受订阅
)

// NotifyMessage 微信小程序推送的通知消息
type NotifyMessage struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   int    `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	List         any    `json:"List"`
	Encrypt      string `json:"Encrypt"`
}

// SubscribeMsgPopupEventListItem 订阅消息弹框事件的结果列表
type SubscribeMsgPopupEventListItem struct {
	PopupScene            string `json:"PopupScene"`
	SubscribeStatusString string `json:"SubscribeStatusString"`
	TemplateId            string `json:"TemplateId"`
}

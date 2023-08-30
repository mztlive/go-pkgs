package officialaccount

import (
	"context"
	"encoding/xml"
	"net/http"
)

// MsgType 消息类型
type MsgType string

// EventType 事件类型
type EventType string

// InfoType 第三方平台授权事件类型
type InfoType string

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName" json:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName" json:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime" json:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType" json:"MsgType"`
}

// SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName CDATA) {
	msg.ToUserName = toUserName
}

// SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName CDATA) {
	msg.FromUserName = fromUserName
}

// SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

// SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}

// GetOpenID get the FromUserName value
func (msg *CommonToken) GetOpenID() string {
	return string(msg.FromUserName)
}

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

// EventPic 发图事件推送
type EventPic struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

// SubscribeMsgPopupEvent 订阅通知事件推送的消息体
type SubscribeMsgPopupEvent struct {
	TemplateID            string `xml:"TemplateId" json:"TemplateId"`
	SubscribeStatusString string `xml:"SubscribeStatusString" json:"SubscribeStatusString"`
	PopupScene            int    `xml:"PopupScene" json:"PopupScene,string"`
}

// MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
	CommonToken

	// 基本消息
	MsgID         int64   `xml:"MsgId"` // 其他消息推送过来是MsgId
	TemplateMsgID int64   `xml:"MsgID"` // 模板消息推送成功的消息是MsgID
	Content       string  `xml:"Content"`
	Recognition   string  `xml:"Recognition"`
	PicURL        string  `xml:"PicUrl"`
	MediaID       string  `xml:"MediaId"`
	Format        string  `xml:"Format"`
	ThumbMediaID  string  `xml:"ThumbMediaId"`
	LocationX     float64 `xml:"Location_X"`
	LocationY     float64 `xml:"Location_Y"`
	Scale         float64 `xml:"Scale"`
	Label         string  `xml:"Label"`
	Title         string  `xml:"Title"`
	Description   string  `xml:"Description"`
	URL           string  `xml:"Url"`

	// 事件相关
	Event       EventType `xml:"Event" json:"Event"`
	EventKey    string    `xml:"EventKey"`
	Ticket      string    `xml:"Ticket"`
	Latitude    string    `xml:"Latitude"`
	Longitude   string    `xml:"Longitude"`
	Precision   string    `xml:"Precision"`
	MenuID      string    `xml:"MenuId"`
	Status      string    `xml:"Status"`
	SessionFrom string    `xml:"SessionFrom"`
	TotalCount  int64     `xml:"TotalCount"`
	FilterCount int64     `xml:"FilterCount"`
	SentCount   int64     `xml:"SentCount"`
	ErrorCount  int64     `xml:"ErrorCount"`

	ScanCodeInfo struct {
		ScanType   string `xml:"ScanType"`
		ScanResult string `xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	SendPicsInfo struct {
		Count   int32      `xml:"Count"`
		PicList []EventPic `xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX float64 `xml:"Location_X"`
		LocationY float64 `xml:"Location_Y"`
		Scale     float64 `xml:"Scale"`
		Label     string  `xml:"Label"`
		Poiname   string  `xml:"Poiname"`
	}

	subscribeMsgPopupEventList []SubscribeMsgPopupEvent `json:"-"`

	SubscribeMsgPopupEvent []struct {
		List SubscribeMsgPopupEvent `xml:"List"`
	} `xml:"SubscribeMsgPopupEvent"`

	//小程序审核通知
	SuccTime   int    `xml:"SuccTime"`   //审核成功时的时间戳
	FailTime   int    `xml:"FailTime"`   //审核不通过的时间戳
	DelayTime  int    `xml:"DelayTime"`  //审核延后时的时间戳
	Reason     string `xml:"Reason"`     //审核不通过的原因
	ScreenShot string `xml:"ScreenShot"` //审核不通过的截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

const (
	// MsgTypeText 表示文本消息
	MsgTypeText MsgType = "text"
	// MsgTypeImage 表示图片消息
	MsgTypeImage MsgType = "image"
	// MsgTypeVoice 表示语音消息
	MsgTypeVoice MsgType = "voice"
	// MsgTypeVideo 表示视频消息
	MsgTypeVideo MsgType = "video"
	// MsgTypeMiniprogrampage 表示小程序卡片消息
	MsgTypeMiniprogrampage MsgType = "miniprogrampage"
	// MsgTypeShortVideo 表示短视频消息[限接收]
	MsgTypeShortVideo MsgType = "shortvideo"
	// MsgTypeLocation 表示坐标消息[限接收]
	MsgTypeLocation MsgType = "location"
	// MsgTypeLink 表示链接消息[限接收]
	MsgTypeLink MsgType = "link"
	// MsgTypeMusic 表示音乐消息[限回复]
	MsgTypeMusic MsgType = "music"
	// MsgTypeNews 表示图文消息[限回复]
	MsgTypeNews MsgType = "news"
	// MsgTypeTransfer 表示消息消息转发到客服
	MsgTypeTransfer MsgType = "transfer_customer_service"
	// MsgTypeEvent 表示事件推送消息
	MsgTypeEvent MsgType = "event"
)

const (
	// EventSubscribe 订阅
	EventSubscribe EventType = "subscribe"
	// EventUnsubscribe 取消订阅
	EventUnsubscribe EventType = "unsubscribe"
	// EventScan 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventScan EventType = "SCAN"
	// EventLocation 上报地理位置事件
	EventLocation EventType = "LOCATION"
	// EventClick 点击菜单拉取消息时的事件推送
	EventClick EventType = "CLICK"
	// EventView 点击菜单跳转链接时的事件推送
	EventView EventType = "VIEW"
	// EventScancodePush 扫码推事件的事件推送
	EventScancodePush EventType = "scancode_push"
	// EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventScancodeWaitmsg EventType = "scancode_waitmsg"
	// EventPicSysphoto 弹出系统拍照发图的事件推送
	EventPicSysphoto EventType = "pic_sysphoto"
	// EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
	EventPicPhotoOrAlbum EventType = "pic_photo_or_album"
	// EventPicWeixin 弹出微信相册发图器的事件推送
	EventPicWeixin EventType = "pic_weixin"
	// EventLocationSelect 弹出地理位置选择器的事件推送
	EventLocationSelect EventType = "location_select"
	// EventViewMiniprogram 点击菜单跳转小程序的事件推送
	EventViewMiniprogram EventType = "view_miniprogram"
	// EventTemplateSendJobFinish 发送模板消息推送通知
	EventTemplateSendJobFinish EventType = "TEMPLATESENDJOBFINISH"
	// EventMassSendJobFinish 群发消息推送通知
	EventMassSendJobFinish EventType = "MASSSENDJOBFINISH"
	// EventWxaMediaCheck 异步校验图片/音频是否含有违法违规内容推送事件
	EventWxaMediaCheck EventType = "wxa_media_check"
	// EventSubscribeMsgPopupEvent 订阅通知事件推送
	EventSubscribeMsgPopupEvent EventType = "subscribe_msg_popup_event"
	// EventPublishJobFinish 发布任务完成
	EventPublishJobFinish EventType = "PUBLISHJOBFINISH"
	// EventWeappAuditSuccess 审核通过
	EventWeappAuditSuccess EventType = "weapp_audit_success"
	// EventWeappAuditFail 审核不通过
	EventWeappAuditFail EventType = "weapp_audit_fail"
	// EventWeappAuditDelay 审核延后
	EventWeappAuditDelay EventType = "weapp_audit_delay"
)

const (
	// 微信开放平台需要用到

	// InfoTypeVerifyTicket 返回ticket
	InfoTypeVerifyTicket InfoType = "component_verify_ticket"
	// InfoTypeAuthorized 授权
	InfoTypeAuthorized InfoType = "authorized"
	// InfoTypeUnauthorized 取消授权
	InfoTypeUnauthorized InfoType = "unauthorized"
	// InfoTypeUpdateAuthorized 更新授权
	InfoTypeUpdateAuthorized InfoType = "updateauthorized"
	// InfoTypeNotifyThirdFasterRegister 注册审核事件推送
	InfoTypeNotifyThirdFasterRegister InfoType = "notify_third_fasteregister"
)

// Reply 消息回复
type Reply struct {
	MsgType MsgType
	MsgData interface{}
}

type MessageHandlerFunc func(ctx context.Context, msg MixMessage) *Reply

type IOfficialAccountServer interface {

	// SetMessageHandler 设置公众号消息处理器
	SetMessageHandler(handler MessageHandlerFunc)

	// Server 接收并处理微信消息
	Server(ctx context.Context, request *http.Request, resp http.ResponseWriter)
}

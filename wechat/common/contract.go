package wechat

import (
	"context"
	"net/http"
)

const (
	// MessageHandlerIsOfficial 公众号
	MessageHandlerIsOfficial = "official"

	// MessageHandlerIsMini 小程序
	MessageHandlerIsMini = "mini"
)

// MessageHandler 处理公众号、小程序消息推送的接口
type MessageHandler interface {

	// 监听消息, 这个方法应该是阻塞的，直到收到消息或者上下文被取消
	// 调用方应该开启一个协程来调用这个方法
	Listener(ctx context.Context, request http.Request, resp http.ResponseWriter)

	// 检查签名
	// CheckSignature(signature, timestamp, nonce string) bool

	// 获取应用类型 (小程序、公众号)
	AppType() string
}

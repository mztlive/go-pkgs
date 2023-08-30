package wrapper

import (
	"context"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/mztlive/project-template/pkg/logger"
	"github.com/mztlive/project-template/pkg/wechat/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"go.uber.org/zap"
)

func (m *SilenceperOfficialAccount) Server(ctx context.Context, request *http.Request, resp http.ResponseWriter) {
	server := m.engine.GetServer(request, resp)

	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		var mixMessage = officialaccount.MixMessage{}
		copier.Copy(&mixMessage, &msg)
		reply := m.messageHandler(ctx, mixMessage)
		if reply == nil {
			return nil
		}

		var replyMessage = message.Reply{}
		copier.Copy(&replyMessage, &reply)
		return &replyMessage
	})

	if err := server.Serve(); err != nil {
		logger.Error("can not serve official account message: ", zap.Error(err))
		return
	}

	if err := server.Send(); err != nil {
		logger.Error("can not send official account message: ", zap.Error(err))
	}
}

func (m *SilenceperOfficialAccount) SetMessageHandler(handler officialaccount.MessageHandlerFunc) {
	m.messageHandler = handler
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MessageSearch struct {
	ConversationId string `json:"conversationId" form:"conversationId"`
	Content        string `json:"content" form:"content" `
	request.PageInfo
}

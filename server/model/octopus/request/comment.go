package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CommentSearch struct {
	ConversationId string `json:"conversationId" form:"conversationId"`
	Content        string `json:"content" form:"content" `
	Status         uint   `json:"status" form:"status" `
	request.PageInfo
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CmtConversationSearch struct {
	ThreadId       string `json:"threadId" form:"threadId"`
	Commenter      string `json:"commenter" form:"commenter" `
	CommentReplier string `json:"commentReplier" form:"commentReplier" `
	request.PageInfo
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CmtConversationSearch struct {
	ThreadId       uint   `json:"threadId"`
	Commenter      string `json:"commenter" form:"commenter" `
	CommentReplier string `json:"commentReplier" form:"commentReplier" `
	request.PageInfo
}

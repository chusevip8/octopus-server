package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CmtConversationSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Commenter  string `json:"commenter" form:"commenter" `
    CommentReplier  string `json:"commentReplier" form:"commentReplier" `
    request.PageInfo
}

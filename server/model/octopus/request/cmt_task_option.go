package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CmtTaskOptionSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    TaskTitle  string `json:"taskTitle" form:"taskTitle" `
    Keyword  string `json:"keyword" form:"keyword" `
    Commenter  string `json:"commenter" form:"commenter" `
    CommenterId  string `json:"commenterId" form:"commenterId" `
    request.PageInfo
}

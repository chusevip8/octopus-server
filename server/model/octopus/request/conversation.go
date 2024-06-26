package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ConversationSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    AppName  string `json:"appName" form:"appName" `
    TaskType  string `json:"taskType" form:"taskType" `
    request.PageInfo
}

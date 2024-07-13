package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CmtTaskSetupSearch struct {
	AppName   string `json:"appName" form:"appName"`
	Keyword   string `json:"keyword" form:"keyword" `
	TaskTitle string `json:"taskTitle" form:"taskTitle"`
	request.PageInfo
}

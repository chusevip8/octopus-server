package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type MsgConversationSearch struct {
	AppName   string `json:"appName" form:"appName"`
	Sender    string `json:"sender" form:"sender" `
	Receiver  string `json:"receiver" form:"receiver" `
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	request.PageInfo
}

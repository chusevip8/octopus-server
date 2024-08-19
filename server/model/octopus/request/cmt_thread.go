package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CmtThreadSearch struct {
	AppName   string `json:"appName" form:"appName"`
	Poster    string `json:"poster" form:"poster" `
	PostTitle string `json:"postTitle" form:"postTitle" `
	PostDesc  string `json:"postDesc" form:"postDesc" `
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	request.PageInfo
}

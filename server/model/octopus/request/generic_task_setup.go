package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GenericTaskSetupSearch struct {
	TaskTitle string `json:"taskTitle" form:"taskTitle"`
	CreatedBy uint   `json:"createdBy" form:"createdBy"`
	request.PageInfo
}

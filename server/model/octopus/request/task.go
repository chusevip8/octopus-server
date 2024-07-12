package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TaskSearch struct {
	TaskSetupId  string `json:"taskSetupId" form:"taskSetupId"`
	AppName      string `json:"appName" form:"appName"`
	Type         string `json:"type" form:"type"`
	DeviceNumber string `json:"deviceNumber" form:"deviceNumber"`
	Status       uint   `json:"status" form:"status" `
	request.PageInfo
}

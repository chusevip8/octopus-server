package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TaskSearch struct {
	TaskSetupId  string `json:"taskSetupId" form:"taskSetupId"`
	MainTaskType string `json:"mainTaskType" form:"mainTaskType"`
	AppName      string `json:"appName" form:"appName"`
	Type         string `json:"type" form:"type"`
	DeviceNumber string `json:"deviceNumber" form:"deviceNumber"`
	Status       uint   `json:"status" form:"status" `
	CreatedBy    uint   `json:"createdBy" form:"createdBy"`
	request.PageInfo
}

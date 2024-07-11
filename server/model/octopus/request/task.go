package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TaskSearch struct {
	AppName      string `json:"appName"`
	Type         string `json:"type"`
	DeviceNumber string `json:"deviceNumber"`
	Status       uint   `json:"status" form:"status" `
	request.PageInfo
}

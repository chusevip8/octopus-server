package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DeviceSearch struct {
	NickName  string `json:"nickName" form:"nickName"`
	Number    string `json:"number" form:"number" `
	Note      string `json:"note" form:"note" `
	Status    int    `json:"status" form:"status"`
	Group     string `json:"group" form:"group"`
	CreatedBy uint   `json:"createdBy" form:"createdBy"`
	request.PageInfo
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CommentSearch struct {
	Content string `json:"content" form:"content" `
	Status  *int   `json:"status" form:"status" `
	request.PageInfo
}

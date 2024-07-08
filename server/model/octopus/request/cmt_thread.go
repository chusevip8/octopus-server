package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CmtThreadSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Poster  string `json:"poster" form:"poster" `
    PostTitle  string `json:"postTitle" form:"postTitle" `
    PostDesc  string `json:"postDesc" form:"postDesc" `
    request.PageInfo
}

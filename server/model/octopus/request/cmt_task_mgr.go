package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CmtTaskMgrSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      TaskTitle  string `json:"taskTitle" form:"taskTitle" `
                      ArticleID  string `json:"articleID" form:"articleID" `
                      CmtKeyword  string `json:"cmtKeyword" form:"cmtKeyword" `
    request.PageInfo
}

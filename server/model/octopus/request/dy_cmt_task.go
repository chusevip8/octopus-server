package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type DYCmtTaskSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      VideoTitle  string `json:"videoTitle" form:"videoTitle" `
                      Keyword  string `json:"keyword" form:"keyword" `
    request.PageInfo
}

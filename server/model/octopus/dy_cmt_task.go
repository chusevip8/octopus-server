// 自动生成模板DYCmtTask
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 抖音评论任务 结构体  DYCmtTask
type DYCmtTask struct {
 global.GVA_MODEL 
      VideoTitle  string `json:"videoTitle" form:"videoTitle" gorm:"column:video_title;comment:视频标题;size:128;" binding:"required"`  //视频标题 
      VideoId  string `json:"videoId" form:"videoId" gorm:"column:video_id;comment:视频ID;" binding:"required"`  //视频ID 
      Keyword  string `json:"keyword" form:"keyword" gorm:"column:keyword;comment:关键字;"`  //关键字 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 抖音评论任务 DYCmtTask自定义表名 oct_dy_cmt_task
func (DYCmtTask) TableName() string {
  return "oct_dy_cmt_task"
}


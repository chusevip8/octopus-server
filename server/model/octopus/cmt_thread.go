// 自动生成模板CmtThread
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论会话 结构体  CmtThread
type CmtThread struct {
	global.GVA_MODEL
	AppName      string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;size:20;" binding:"required"`           //应用名称
	TaskOptionId uint   `json:"taskOptionId" form:"taskOptionId" gorm:"column:task_option_id;comment:任务参数Id;" binding:"required"` //任务参数Id
	PostId       string `json:"postId" form:"postId" gorm:"column:post_id;comment:帖子Id;" binding:"required"`                      //帖子Id
	Poster       string `json:"poster" form:"poster" gorm:"column:poster;comment:发帖者;size:64;"`                                   //发帖者
	PostTitle    string `json:"postTitle" form:"postTitle" gorm:"column:post_title;comment:帖子标题;"`                                //帖子标题
	PostDesc     string `json:"postDesc" form:"postDesc" gorm:"column:post_desc;comment:帖子描述;"`                                   //帖子描述
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 评论会话 CmtThread自定义表名 oct_cmt_thread
func (CmtThread) TableName() string {
	return "oct_cmt_thread"
}

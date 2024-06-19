// 自动生成模板CommentTask
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论任务 结构体  CommentTask
type CommentTask struct {
	global.GVA_MODEL
	AppName   string `json:"appName" form:"appName" gorm:"column:app_name;comment:App名称;size:20;" binding:"required"` //App名称
	Title     string `json:"title" form:"title" gorm:"column:title;comment:任务标题;size:128;" binding:"required"`        //任务标题
	ScriptID  uint   `json:"scriptID" form:"scriptID" gorm:"column:script_id;comment:脚本ID;" binding:"required"`       //脚本ID
	ArticleID string `json:"articleID" form:"articleID" gorm:"column:article_id;comment:文章ID;" binding:"required"`    //文章ID
	Keyword   string `json:"keyword" form:"keyword" gorm:"column:keyword;comment:评论关键字;"`                             //评论关键字
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 评论任务 CommentTask自定义表名 oct_comment_task
func (CommentTask) TableName() string {
	return "oct_comment_task"
}

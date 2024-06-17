// 自动生成模板CmtTaskMgr
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论任务管理 结构体  CmtTaskMgr
type CmtTaskMgr struct {
	global.GVA_MODEL
	AppName    string `json:"appName" form:"appName" gorm:"column:app_name;comment:App名称;size:20;" binding:"required"`       //App名称
	TaskTitle  string `json:"taskTitle" form:"taskTitle" gorm:"column:task_title;comment:任务标题;size:128;" binding:"required"` //任务标题
	ArticleID  string `json:"articleID" form:"articleID" gorm:"column:article_id;comment:文章ID;" binding:"required"`          //文章ID
	CmtKeyword string `json:"cmtKeyword" form:"cmtKeyword" gorm:"column:cmt_keyword;comment:评论关键字;"`                         //评论关键字
	ScriptID   uint   `json:"scriptID" form:"scriptID" gorm:"column:script_id;comment:脚本ID;" binding:"required"`             //脚本ID
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 评论任务管理 CmtTaskMgr自定义表名 oct_cmt_task_mgr
func (CmtTaskMgr) TableName() string {
	return "oct_cmt_task_mgr"
}

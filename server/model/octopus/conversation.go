// 自动生成模板Conversation
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 消息会话 结构体  Conversation
type Conversation struct {
	global.GVA_MODEL
	AppName     string      `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;size:20;" binding:"required"` //App名称
	TaskID      uint        `json:"taskID" form:"taskID" gorm:"column:task_id;comment:任务ID;" binding:"required"`            //任务ID
	CommentTask CommentTask `json:"commentTask" gorm:"foreignKey:TaskID;references:ID;comment:关联评论任务;"`
	CreatedBy   uint        `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint        `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint        `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 消息会话 Conversation自定义表名 oct_conversation
func (Conversation) TableName() string {
	return "oct_conversation"
}

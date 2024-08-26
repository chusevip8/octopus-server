// 自动生成模板MsgTaskSetup
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 私信任务设置 结构体  MsgTaskSetup
type MsgTaskSetup struct {
	global.GVA_MODEL
	AppName   string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;size:20;" binding:"required"` //应用名称
	ScriptId  uint   `json:"scriptId" form:"scriptId" gorm:"column:script_id;comment:脚本Id;" binding:"required"`      //脚本Id
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 私信任务设置 MsgTaskSetup自定义表名 oct_msg_task_setup
func (MsgTaskSetup) TableName() string {
	return "oct_msg_task_setup"
}

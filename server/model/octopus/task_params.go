// 自动生成模板TaskParams
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 任务参数 结构体  TaskParams
type TaskParams struct {
	global.GVA_MODEL
	TaskSetupId uint   `json:"taskSetupId" form:"taskSetupId" gorm:"column:task_setup_id;comment:任务设置Id;" binding:"required"` //任务设置Id
	TaskType    string `json:"taskType" form:"taskType" gorm:"column:task_type;comment:任务类型;" binding:"required"`             //任务类型
	Params      string `json:"params" form:"params" gorm:"column:params;comment:任务参数;" binding:"required"`                    //任务参数
	ScriptId    uint   `json:"scriptId" form:"scriptId" gorm:"column:script_id;comment:脚本Id;" binding:"required"`             //脚本Id
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 任务参数 TaskParams自定义表名 oct_task_params
func (TaskParams) TableName() string {
	return "oct_task_params"
}

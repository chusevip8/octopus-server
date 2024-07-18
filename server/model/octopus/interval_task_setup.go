// 自动生成模板IntervalTaskSetup
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 间隔任务设置 结构体  IntervalTaskSetup
type IntervalTaskSetup struct {
	global.GVA_MODEL
	AppName     string `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;size:20;" binding:"required"`        //应用名称
	TaskTitle   string `json:"taskTitle" form:"taskTitle" gorm:"column:task_title;comment:任务标题;size:128;" binding:"required"` //任务标题
	ScriptId    uint   `json:"scriptId" form:"scriptId" gorm:"column:script_id;comment:脚本Id;" binding:"required"`             //脚本Id
	Params      string `json:"params" form:"params" gorm:"column:params;comment:脚本参数;type:text;"`                             //脚本参数
	IntervalMin uint   `json:"intervalMin" form:"intervalMin" gorm:"column:interval_min;comment:间隔时间;" binding:"required"`    //间隔时间
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 间隔任务设置 IntervalTaskSetup自定义表名 oct_interval_task_setup
func (IntervalTaskSetup) TableName() string {
	return "oct_interval_task_setup"
}

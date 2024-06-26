// 自动生成模板ExecTask
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 执行任务 结构体  ExecTask
type ExecTask struct {
	global.GVA_MODEL
	AppName   string `json:"appName" form:"appName" gorm:"column:app_name;comment:App名称;size:20;" binding:"required"`   //App名称
	TaskType  string `json:"taskType" form:"taskType" gorm:"column:task_type;comment:任务类型;size:20;" binding:"required"` //任务类型
	TaskID    uint   `json:"taskID" form:"taskID" gorm:"column:task_id;comment:任务ID;" binding:"required"`               //任务ID
	DeviceID  uint   `json:"deviceID" form:"deviceID" gorm:"column:device_id;comment:设备ID;" binding:"required"`         //设备ID
	Device    Device `json:"device" gorm:"foreignKey:DeviceID;references:ID;comment:关联设备;"`
	Status    uint   `json:"status" form:"status" gorm:"column:status;comment:任务状态;" binding:"required"` //任务状态
	Error     string `json:"error" form:"error" gorm:"column:error;comment:任务错误信息;"`                     //任务错误信息
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 执行任务 ExecTask自定义表名 exec_task
func (ExecTask) TableName() string {
	return "oct_exec_task"
}

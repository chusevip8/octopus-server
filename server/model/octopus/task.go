// 自动生成模板Task
package octopus

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// 任务 结构体  Task
type Task struct {
	global.GVA_MODEL
	AppName      string     `json:"appName" form:"appName" gorm:"column:app_name;comment:App名称;size:20;" binding:"required"`          //App名称
	TaskParamsId uint       `json:"taskParamsId" form:"taskParamsId" gorm:"column:task_params_id;comment:任务参数Id;" binding:"required"` //任务参数Id
	TaskParams   TaskParams `json:"taskParams" gorm:"foreignKey:TaskParamsId;references:ID;comment:关联任务参数;"`
	DeviceId     uint       `json:"deviceId" form:"deviceId" gorm:"column:device_id;comment:设备Id;" binding:"required"` //设备Id
	Device       Device     `json:"device" gorm:"foreignKey:DeviceId;references:ID;comment:关联设备;"`
	Status       uint       `json:"status" form:"status" gorm:"column:status;comment:任务状态;" binding:"required"` //任务状态
	Error        string     `json:"error" form:"error" gorm:"column:error;comment:任务错误信息;"`                     //任务错误信息
	CreatedBy    uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 任务 Task自定义表名 oct_task
func (Task) TableName() string {
	return "oct_task"
}

func (task *Task) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Printf("New user created:(ID: %d)\n", task.ID)
	return
}

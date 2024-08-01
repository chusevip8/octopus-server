// 自动生成模板GenericTaskSetup
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 通用任务设置 结构体  GenericTaskSetup
type GenericTaskSetup struct {
	global.GVA_MODEL
	AppName   string     `json:"appName" form:"appName" gorm:"column:app_name;comment:应用名称;size:20;" binding:"required"`        //应用名称
	TaskTitle string     `json:"taskTitle" form:"taskTitle" gorm:"column:task_title;comment:任务标题;size:128;" binding:"required"` //任务标题
	ScriptId  uint       `json:"scriptId" form:"scriptId" gorm:"column:script_id;comment:脚本Id;" binding:"required"`             //脚本Id
	StartAt   *time.Time `json:"startAt" form:"startAt" gorm:"column:start_at;comment:启动时间;"`                                   //启动时间
	DataFile  string     `json:"dataFile" form:"dataFile" gorm:"column:data_file;comment:数据文件;size:128;"`                       //数据文件
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 通用任务设置 GenericTaskSetup自定义表名 oct_generic_task_setup
func (GenericTaskSetup) TableName() string {
	return "oct_generic_task_setup"
}

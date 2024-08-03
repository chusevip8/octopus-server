// 自动生成模板TaskBindData
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 任务数据 结构体  TaskBindData
type TaskBindData struct {
	global.GVA_MODEL
	TaskSetupId  uint   `json:"taskSetupId" form:"taskSetupId" gorm:"column:task_setup_id;comment:任务设置Id;" binding:"required"`   //任务设置Id
	MainTaskType string `json:"mainTaskType" form:"mainTaskType" gorm:"column:main_task_type;comment:主任务类型;" binding:"required"` //主任务类型
	Status       uint   `json:"status" form:"status" gorm:"default:1;column:status;comment:Status;"`                             //状态
	Item1        string `json:"item1" form:"item1" gorm:"column:item1;comment:字段1;"`                                             //字段1
	Item2        string `json:"item2" form:"item2" gorm:"column:item2;comment:字段2;"`                                             //字段2
	Item3        string `json:"item3" form:"item3" gorm:"column:item3;comment:字段3;"`                                             //字段3
	Item4        string `json:"item4" form:"item4" gorm:"column:item4;comment:字段4;"`                                             //字段4
	Item5        string `json:"item5" form:"item5" gorm:"column:item5;comment:字段5;"`                                             //字段5
	Item6        string `json:"item6" form:"item6" gorm:"column:item6;comment:字段6;"`                                             //字段6
	Item7        string `json:"item7" form:"item7" gorm:"column:item7;comment:字段7;"`                                             //字段7
	Item8        string `json:"item8" form:"item8" gorm:"column:item8;comment:字段8;"`                                             //字段8
	Item9        string `json:"item9" form:"item9" gorm:"column:item9;comment:字段9;"`                                             //字段9
	Item10       string `json:"item10" form:"item10" gorm:"column:item10;comment:字段10;"`                                         //字段10
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 任务数据 TaskBindData自定义表名 oct_task_bind_data
func (TaskBindData) TableName() string {
	return "oct_task_bind_data"
}

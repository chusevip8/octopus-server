// 自动生成模板CmtTaskParams
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 评论任务参数 结构体  CmtTaskParams
type CmtTaskParams struct {
	global.GVA_MODEL
	SetupId   uint   `json:"setupId" form:"setupId" gorm:"column:setup_id;comment:任务设置Id;" binding:"required"`  //任务设置Id
	Params    string `json:"params" form:"params" gorm:"column:params;comment:任务参数;" binding:"required"`        //任务参数
	ScriptId  uint   `json:"scriptId" form:"scriptId" gorm:"column:script_id;comment:脚本Id;" binding:"required"` //脚本Id
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 评论任务参数 CmtTaskParams自定义表名 oct_cmt_task_params
func (CmtTaskParams) TableName() string {
	return "oct_cmt_task_params"
}

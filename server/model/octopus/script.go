// 自动生成模板Script
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 脚本 结构体  Script
type Script struct {
 global.GVA_MODEL 
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:50;" binding:"required"`  //标题 
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;" binding:"required"`  //内容 
      Note  string `json:"note" form:"note" gorm:"column:note;comment:备注;size:128;"`  //备注 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 脚本 Script自定义表名 oct_script
func (Script) TableName() string {
  return "oct_script"
}


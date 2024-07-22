package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type Device struct {
	global.GVA_MODEL
	Number    string         `json:"number" form:"number" gorm:"index;column:number;comment:编号;size:20;not null;"`
	Brand     string         `json:"brand" form:"brand" gorm:"column:brand;comment:品牌;size:50;"`
	OS        string         `json:"os" form:"os" gorm:"column:os;comment:系统;size:50;"`
	Note      string         `json:"note" form:"note" gorm:"column:note;comment:备注;size:64;"`
	Group     string         `json:"group"  form:"group" gorm:"column:group;comment:设备分组;size:10;"`
	Status    int            `json:"status" gorm:"column:status;comment:状态;default:2;"`
	Username  string         `json:"userName" gorm:"index;column:username;comment:关联用户登录名;not null;"`
	User      system.SysUser `json:"user" gorm:"foreignKey:Username;references:Username;comment:关联用户;"`
	CreatedBy uint           `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint           `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint           `gorm:"column:deleted_by;comment:删除者"`
}

func (Device) TableName() string {
	return "oct_device"
}

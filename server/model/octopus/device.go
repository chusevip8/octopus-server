package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type Device struct {
	global.GVA_MODEL
	Number   string         `json:"number" form:"number" gorm:"column:number;comment:编号;size:20;"`
	Brand    string         `json:"brand" form:"brand" gorm:"column:brand;comment:品牌;size:50;"`
	OS       string         `json:"os" form:"os" gorm:"column:os;comment:系统;size:50;"`
	Note     string         `json:"note" form:"note" gorm:"column:note;comment:备注;size:64;"`
	Status   int            `json:"status" gorm:"column:status;comment:状态;default:4;"`
	Username string         `json:"userName" gorm:"index;column:username;comment:关联用户账号"`
	User     system.SysUser `json:"user" gorm:"foreignKey:Username;references:Username;comment:关联用户信息"`
}

func (Device) TableName() string {
	return "oct_device"
}

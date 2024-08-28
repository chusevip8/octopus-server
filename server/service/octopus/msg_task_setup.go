package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type MsgTaskSetupService struct{}

var MsgTaskSetupServiceApp = MsgTaskSetupService{}

// CreateMsgTaskSetup 创建私信任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgTaskSetupService *MsgTaskSetupService) CreateMsgTaskSetup(msgTaskSetup *octopus.MsgTaskSetup) (err error) {
	err = global.GVA_DB.Create(msgTaskSetup).Error
	return err
}

// DeleteMsgTaskSetup 删除私信任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgTaskSetupService *MsgTaskSetupService) DeleteMsgTaskSetup(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.MsgTaskSetup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.MsgTaskSetup{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteMsgTaskSetupByIds 批量删除私信任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgTaskSetupService *MsgTaskSetupService) DeleteMsgTaskSetupByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.MsgTaskSetup{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.MsgTaskSetup{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateMsgTaskSetup 更新私信任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgTaskSetupService *MsgTaskSetupService) UpdateMsgTaskSetup(msgTaskSetup octopus.MsgTaskSetup) (err error) {
	err = global.GVA_DB.Model(&octopus.MsgTaskSetup{}).Where("id = ?", msgTaskSetup.ID).Updates(&msgTaskSetup).Error
	return err
}

// GetMsgTaskSetup 根据ID获取私信任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgTaskSetupService *MsgTaskSetupService) GetMsgTaskSetup(ID string) (msgTaskSetup octopus.MsgTaskSetup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&msgTaskSetup).Error
	return
}

func (msgTaskSetupService *MsgTaskSetupService) GetMsgTaskSetupByAppName(appName string) (msgTaskSetup octopus.MsgTaskSetup, err error) {
	err = global.GVA_DB.Where("app_name = ?", appName).First(&msgTaskSetup).Error
	return
}

// GetMsgTaskSetupInfoList 分页获取私信任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (msgTaskSetupService *MsgTaskSetupService) GetMsgTaskSetupInfoList(info octopusReq.MsgTaskSetupSearch) (list []octopus.MsgTaskSetup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.MsgTaskSetup{})
	var msgTaskSetups []octopus.MsgTaskSetup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&msgTaskSetups).Error
	return msgTaskSetups, total, err
}

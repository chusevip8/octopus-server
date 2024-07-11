package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type CmtTaskSetupService struct{}

var CmtTaskSetupServiceApp = new(CmtTaskSetupService)

// CreateCmtTaskSetup 创建评论任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskSetupService *CmtTaskSetupService) CreateCmtTaskSetup(cmtTaskSetup *octopus.CmtTaskSetup) (err error) {
	err = global.GVA_DB.Create(cmtTaskSetup).Error
	return err
}

// DeleteCmtTaskSetup 删除评论任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskSetupService *CmtTaskSetupService) DeleteCmtTaskSetup(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtTaskSetup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.CmtTaskSetup{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCmtTaskSetupByIds 批量删除评论任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskSetupService *CmtTaskSetupService) DeleteCmtTaskSetupByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.CmtTaskSetup{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.CmtTaskSetup{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCmtTaskSetup 更新评论任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskSetupService *CmtTaskSetupService) UpdateCmtTaskSetup(cmtTaskSetup octopus.CmtTaskSetup) (err error) {
	err = global.GVA_DB.Model(&octopus.CmtTaskSetup{}).Where("id = ?", cmtTaskSetup.ID).Updates(&cmtTaskSetup).Error
	return err
}

// GetCmtTaskSetup 根据ID获取评论任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskSetupService *CmtTaskSetupService) GetCmtTaskSetup(ID string) (cmtTaskSetup octopus.CmtTaskSetup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&cmtTaskSetup).Error
	return
}

// GetCmtTaskSetupInfoList 分页获取评论任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmtTaskSetupService *CmtTaskSetupService) GetCmtTaskSetupInfoList(info octopusReq.CmtTaskSetupSearch) (list []octopus.CmtTaskSetup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.CmtTaskSetup{})
	var cmtTaskSetups []octopus.CmtTaskSetup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Keyword != "" {
		db = db.Where("keyword LIKE ?", "%"+info.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&cmtTaskSetups).Error
	return cmtTaskSetups, total, err
}

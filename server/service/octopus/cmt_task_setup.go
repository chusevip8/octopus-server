package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
	"time"
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
//func (cmtTaskSetupService *CmtTaskSetupService) DeleteCmtTaskSetup(ID string, userID uint) (err error) {
//	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
//		if err := tx.Model(&octopus.CmtTaskSetup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
//			return err
//		}
//		if err = tx.Delete(&octopus.CmtTaskSetup{}, "id = ?", ID).Error; err != nil {
//			return err
//		}
//		return nil
//	})
//	return err
//}

func (cmtTaskSetupService *CmtTaskSetupService) DeleteCmtTaskSetup(ID string, userID uint) (err error) {
	var taskIds []uint
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ?", ID, "cmt").
		Pluck("oct_task.id", &taskIds).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		if err := tx.Model(&octopus.Task{}).Where("id in ?", taskIds).Updates(map[string]interface{}{
			"deleted_by": userID,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&octopus.TaskParams{}).Where("task_setup_id = ?", ID).Where("main_task_type = ?", "cmt").Updates(map[string]interface{}{
			"deleted_by": userID,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&octopus.CmtTaskSetup{}).Where("id = ?", ID).Updates(map[string]interface{}{
			"deleted_by": userID,
			"deleted_at": now,
		}).Error; err != nil {
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
	err = CmtTaskServiceApp.UpdateReadPostCmtTaskParams(cmtTaskSetup)
	if err == nil {
		err = global.GVA_DB.Model(&octopus.CmtTaskSetup{}).Where("id = ?", cmtTaskSetup.ID).Updates(&cmtTaskSetup).Error
	}
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

	filter(db, info.CreatedBy)

	var cmtTaskSetups []octopus.CmtTaskSetup
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.AppName != "" {
		db = db.Where("app_name = ?", info.AppName)
	}
	if info.TaskTitle != "" {
		db = db.Where("task_title LIKE ?", "%"+info.Keyword+"%")
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

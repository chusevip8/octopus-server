package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
	"os"
)

type GenericTaskSetupService struct{}

var GenericTaskSetupServiceApp = new(GenericTaskSetupService)

// CreateGenericTaskSetup 创建通用任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (genericTaskSetupService *GenericTaskSetupService) CreateGenericTaskSetup(genericTaskSetup *octopus.GenericTaskSetup) (err error) {
	err = global.GVA_DB.Create(genericTaskSetup).Error
	return err
}

// DeleteGenericTaskSetup 删除通用任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (genericTaskSetupService *GenericTaskSetupService) DeleteGenericTaskSetup(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.GenericTaskSetup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.GenericTaskSetup{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteGenericTaskSetupByIds 批量删除通用任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (genericTaskSetupService *GenericTaskSetupService) DeleteGenericTaskSetupByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.GenericTaskSetup{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.GenericTaskSetup{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateGenericTaskSetup 更新通用任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (genericTaskSetupService *GenericTaskSetupService) UpdateGenericTaskSetup(genericTaskSetup octopus.GenericTaskSetup) (err error) {
	err = global.GVA_DB.Model(&octopus.GenericTaskSetup{}).Where("id = ?", genericTaskSetup.ID).Save(&genericTaskSetup).Error
	return err
}

// GetGenericTaskSetup 根据ID获取通用任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (genericTaskSetupService *GenericTaskSetupService) GetGenericTaskSetup(ID string) (genericTaskSetup octopus.GenericTaskSetup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&genericTaskSetup).Error
	return
}

// GetGenericTaskSetupInfoList 分页获取通用任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (genericTaskSetupService *GenericTaskSetupService) GetGenericTaskSetupInfoList(info octopusReq.GenericTaskSetupSearch) (list []octopus.GenericTaskSetup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.GenericTaskSetup{})
	var genericTaskSetups []octopus.GenericTaskSetup
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TaskTitle != "" {
		db = db.Where("task_title LIKE ?", "%"+info.TaskTitle+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&genericTaskSetups).Error
	return genericTaskSetups, total, err
}

func (genericTaskSetupService *GenericTaskSetupService) DeleteBindData(setupId string, mainTaskType string) (err error) {
	taskSetup, err := genericTaskSetupService.GetGenericTaskSetup(setupId)
	if err != nil {
		return err
	}
	err = taskBindDataServiceApp.DeleteTaskBindDataBySetupId(setupId, mainTaskType)
	if err != nil {
		return err
	}
	err = os.Remove(taskSetup.DataFilePath)
	if err != nil {
		return err
	}
	taskSetup.DataFile = ""
	taskSetup.DataFilePath = ""
	err = genericTaskSetupService.UpdateGenericTaskSetup(taskSetup)
	return err
}

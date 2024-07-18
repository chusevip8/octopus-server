package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
    octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
    "gorm.io/gorm"
)

type IntervalTaskSetupService struct {}

// CreateIntervalTaskSetup 创建间隔任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (intervalTaskSetupService *IntervalTaskSetupService) CreateIntervalTaskSetup(intervalTaskSetup *octopus.IntervalTaskSetup) (err error) {
	err = global.GVA_DB.Create(intervalTaskSetup).Error
	return err
}

// DeleteIntervalTaskSetup 删除间隔任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (intervalTaskSetupService *IntervalTaskSetupService)DeleteIntervalTaskSetup(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.IntervalTaskSetup{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&octopus.IntervalTaskSetup{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteIntervalTaskSetupByIds 批量删除间隔任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (intervalTaskSetupService *IntervalTaskSetupService)DeleteIntervalTaskSetupByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&octopus.IntervalTaskSetup{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&octopus.IntervalTaskSetup{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateIntervalTaskSetup 更新间隔任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (intervalTaskSetupService *IntervalTaskSetupService)UpdateIntervalTaskSetup(intervalTaskSetup octopus.IntervalTaskSetup) (err error) {
	err = global.GVA_DB.Model(&octopus.IntervalTaskSetup{}).Where("id = ?",intervalTaskSetup.ID).Updates(&intervalTaskSetup).Error
	return err
}

// GetIntervalTaskSetup 根据ID获取间隔任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (intervalTaskSetupService *IntervalTaskSetupService)GetIntervalTaskSetup(ID string) (intervalTaskSetup octopus.IntervalTaskSetup, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&intervalTaskSetup).Error
	return
}

// GetIntervalTaskSetupInfoList 分页获取间隔任务设置记录
// Author [piexlmax](https://github.com/piexlmax)
func (intervalTaskSetupService *IntervalTaskSetupService)GetIntervalTaskSetupInfoList(info octopusReq.IntervalTaskSetupSearch) (list []octopus.IntervalTaskSetup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&octopus.IntervalTaskSetup{})
    var intervalTaskSetups []octopus.IntervalTaskSetup
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TaskTitle != "" {
        db = db.Where("task_title LIKE ?","%"+ info.TaskTitle+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&intervalTaskSetups).Error
	return  intervalTaskSetups, total, err
}
package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type ExecTaskService struct {
}

// CreateExecTask 创建执行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (execTaskService *ExecTaskService) CreateExecTask(execTask *octopus.ExecTask) (err error) {
	err = global.GVA_DB.Create(execTask).Error
	return err
}

// DeleteExecTask 删除执行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (execTaskService *ExecTaskService) DeleteExecTask(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.ExecTask{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.ExecTask{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteExecTaskByIds 批量删除执行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (execTaskService *ExecTaskService) DeleteExecTaskByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.ExecTask{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.ExecTask{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateExecTask 更新执行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (execTaskService *ExecTaskService) UpdateExecTask(execTask octopus.ExecTask) (err error) {
	err = global.GVA_DB.Model(&octopus.ExecTask{}).Where("id = ?", execTask.ID).Updates(&execTask).Error
	return err
}

// GetExecTask 根据ID获取执行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (execTaskService *ExecTaskService) GetExecTask(ID string) (execTask octopus.ExecTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&execTask).Error
	return
}

func (execTaskService *ExecTaskService) GetExecTaskByDeviceID(ID string) (execTask octopus.ExecTask, err error) {
	err = global.GVA_DB.Where("device_id = ?", ID).First(&execTask).Error
	return
}

// GetExecTaskInfoList 分页获取执行任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (execTaskService *ExecTaskService) GetExecTaskInfoList(info octopusReq.ExecTaskSearch) (list []octopus.ExecTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.ExecTask{})
	var execTasks []octopus.ExecTask
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

	err = db.Find(&execTasks).Error
	return execTasks, total, err
}

package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type TaskParamsService struct{}

var TaskParamsServiceApp = new(TaskParamsService)

// CreateTaskParams 创建任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParamsService *TaskParamsService) CreateTaskParams(taskParams *octopus.TaskParams) (err error) {
	err = global.GVA_DB.Create(taskParams).Error
	return err
}

// DeleteTaskParams 删除任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParamsService *TaskParamsService) DeleteTaskParams(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.TaskParams{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.TaskParams{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTaskParamsByIds 批量删除任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParamsService *TaskParamsService) DeleteTaskParamsByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.TaskParams{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.TaskParams{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTaskParams 更新任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParamsService *TaskParamsService) UpdateTaskParams(taskParams octopus.TaskParams) (err error) {
	err = global.GVA_DB.Model(&octopus.TaskParams{}).Where("id = ?", taskParams.ID).Updates(&taskParams).Error
	return err
}

// GetTaskParams 根据ID获取任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParamsService *TaskParamsService) GetTaskParams(ID string) (taskParams octopus.TaskParams, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&taskParams).Error
	return
}

// GetTaskParamsInfoList 分页获取任务参数记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskParamsService *TaskParamsService) GetTaskParamsInfoList(info octopusReq.TaskParamsSearch) (list []octopus.TaskParams, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.TaskParams{})
	var taskParamss []octopus.TaskParams
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Params != "" {
		db = db.Where("params LIKE ?", "%"+info.Params+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&taskParamss).Error
	return taskParamss, total, err
}

package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
)

type TaskService struct{}

var TaskServiceApp = new(TaskService)

// CreateTask 创建任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) CreateTask(task *octopus.Task) (err error) {
	err = global.GVA_DB.Create(task).Error
	return err
}

// DeleteTask 删除任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) DeleteTask(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.Task{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&octopus.Task{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTaskByIds 批量删除任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) DeleteTaskByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.Task{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&octopus.Task{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTask 更新任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) UpdateTask(task octopus.Task) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("id = ?", task.ID).Updates(&task).Error
	return err
}

// GetTask 根据ID获取任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) GetTask(ID string) (task octopus.Task, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&task).Error
	return
}

// GetTaskInfoList 分页获取任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) GetTaskInfoList(info octopusReq.TaskSearch) (list []octopus.Task, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&octopus.Task{})
	var tasks []octopus.Task
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tasks).Error
	return tasks, total, err
}

package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
	"time"
)

type TaskService struct{}

var (
	NewTask        = make(chan *octopus.Task)
	StopTask       = make(chan *octopus.Task)
	TaskServiceApp = new(TaskService)
)

// CreateTask 创建任务记录
// Author [piexlmax](https://github.com/piexlmax)
func (taskService *TaskService) CreateTask(task *octopus.Task) (err error) {
	err = global.GVA_DB.Create(task).Error
	if err == nil {
		NewTask <- task
	}
	return err
}

func (taskService *TaskService) CreateTaskWithoutNotify(task *octopus.Task) (err error) {
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

func (taskService *TaskService) UpdateTaskStatusRunToFailByDeviceId(deviceId uint, error string) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("device_id = ?", deviceId).Where("status = ?", 2).Updates(map[string]interface{}{"status": 4, "finish_at": time.Now(), "error": error}).Error
	return
}
func (taskService *TaskService) UpdateTaskStatusToRun(id string) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("id = ?", id).Updates(map[string]interface{}{"status": 2, "error": ""}).Error
	return
}
func (taskService *TaskService) UpdateTaskStatusToFinish(id string) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("id = ?", id).Updates(map[string]interface{}{"status": 3, "finish_at": time.Now(), "error": ""}).Error
	return
}
func (taskService *TaskService) UpdateTaskStatusToFail(id string, error string) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("id = ?", id).Updates(map[string]interface{}{"status": 4, "finish_at": time.Now(), "error": error}).Error
	return
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
	//db := global.GVA_DB.Model(&octopus.Task{}).Preload("TaskParams", "task_setup_id=?", info.TaskSetupId).Preload("Device")

	db := global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ? AND oct_task_params.sub_task_type = ?", info.TaskSetupId, info.MainTaskType, info.SubTaskType).
		Preload("TaskParams").
		Preload("Device")

	if !isAdmin(info.CreatedBy) {
		db.Where("oct_task.created_by = ?", info.CreatedBy)
	}

	var tasks []octopus.Task
	// 如果有条件搜索 下方会自动创建搜索语句

	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.Status != 0 {
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

func (taskService *TaskService) GetTasksByDeviceId(taskSetupId string, deviceId string, mainTaskType string) (task []octopus.Task, err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ? AND oct_task.device_id = ?", taskSetupId, mainTaskType, deviceId).Find(&task).Error
	return
}

func (taskService *TaskService) GetTasksByTaskSetupId(taskSetupId string, mainTaskType string, subTaskType string) (tasks []octopus.Task, err error) {
	query := global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ?", taskSetupId, mainTaskType)

	if subTaskType != "" {
		query = query.Where("oct_task_params.sub_task_type = ?", subTaskType)
	}
	err = query.Find(&tasks).Error

	return
}

func (taskService *TaskService) FindPushTask(deviceId string) (task octopus.Task, err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Preload("TaskParams").
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.main_task_type != ? AND oct_task.device_id = ? AND oct_task.status = ?", "interval", deviceId, 1).
		First(&task).Error
	return
}

func (taskService *TaskService) UpdateTaskStatusRunToFail() {
	_ = global.GVA_DB.Model(&octopus.Task{}).Where("status = ?", 2).Updates(map[string]interface{}{"status": 4, "finish_at": time.Now(), "error": "服务器重启"}).Error
	return
}

func (taskService *TaskService) FindIntervalTasks() (tasks []octopus.Task, err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Preload("TaskParams").
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.main_task_type = ? AND oct_task.status != ?", "interval", 2).Find(&tasks).Error
	return
}

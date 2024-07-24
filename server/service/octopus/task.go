package octopus

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
//
//	func (taskService *TaskService) DeleteTask(ID string, userID uint) (err error) {
//		err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
//			if err := tx.Model(&octopus.Task{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
//				return err
//			}
//			if err = tx.Delete(&octopus.Task{}, "id = ?", ID).Error; err != nil {
//				return err
//			}
//			return nil
//		})
//		return err
//	}
func (taskService *TaskService) DeleteTask(id string, userId uint) (err error) {
	var task octopus.Task
	err = global.GVA_DB.Preload("TaskParams").First(&task, id).Error
	if err != nil {
		return
	}
	var taskIds []uint
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ? AND oct_task.device_id = ?", task.TaskParams.TaskSetupId, task.TaskParams.MainTaskType, task.DeviceId).
		Pluck("oct_task.id", &taskIds).Error
	if err != nil {
		return
	}

	var taskParamsIds []uint
	err = global.GVA_DB.Model(&octopus.Task{}).
		Where("id in ?", taskIds).
		Pluck("oct_task.task_params_id", &taskParamsIds).Error
	if err != nil {
		return
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&octopus.Task{}).Where("id in ?", taskIds).Update("deleted_by", userId).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", taskIds).Delete(&octopus.Task{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&octopus.TaskParams{}).Where("id in ?", taskParamsIds).Update("deleted_by", userId).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", taskParamsIds).Delete(&octopus.TaskParams{}).Error; err != nil {
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
	err = global.GVA_DB.Model(&octopus.Task{}).Where("device_id = ?", deviceId).Where("status = ?", 2).Updates(map[string]interface{}{"status": 4, "error": error}).Error
	return
}
func (taskService *TaskService) UpdateTaskStatusToRun(id string) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("id = ?", id).Updates(map[string]interface{}{"status": 2, "error": ""}).Error
	return
}
func (taskService *TaskService) UpdateTaskStatusToFinish(id string) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("id = ?", id).Updates(map[string]interface{}{"status": 3, "error": ""}).Error
	return
}
func (taskService *TaskService) UpdateTaskStatusToFail(id string, error string) (err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).Where("id = ?", id).Updates(map[string]interface{}{"status": 4, "error": error}).Error
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
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ? AND oct_task.created_by = ?", info.TaskSetupId, info.MainTaskType, info.CreatedBy).
		Preload("TaskParams").
		Preload("Device")
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
func (taskService *TaskService) GetTaskByDeviceId(taskSetupId string, deviceId string, mainTaskType string) (task octopus.Task, err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task_params.main_task_type = ? AND oct_task.device_id = ?", taskSetupId, mainTaskType, deviceId).First(&task).Error
	return
}

func (taskService *TaskService) findPushTask(deviceId string) (task octopus.Task, err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.main_task_type != ? AND oct_task.device_id = ? AND oct_task.status = ?", "interval", deviceId, 1).
		First(&task).Error
	return
}

func (taskService *TaskService) buildTaskPush(task octopus.Task, taskPush *protocol.TaskPush) (err error) {

	var script octopus.Script
	err = global.GVA_DB.Where("id = ?", task.TaskParams.ScriptId).First(&script).Error
	if err != nil {
		taskPush.Error = "Can't find task script"
		return
	}
	var params map[string]string
	err = json.Unmarshal([]byte(task.TaskParams.Params), &params)
	if err != nil {
		taskPush.Error = "Task params json unmarshal error"
		return
	}
	scriptContent := script.Content
	for key, value := range params {
		placeholder := fmt.Sprintf("${%s}", key)
		scriptContent = strings.ReplaceAll(scriptContent, placeholder, value)
	}
	taskPush.TaskId = strconv.Itoa(int(task.ID))
	taskPush.Script = scriptContent
	taskPush.Error = ""
	return
}

func (taskService *TaskService) PushTask(deviceId string) (taskPush protocol.TaskPush, err error) {
	task, err := taskService.findPushTask(deviceId)
	if err != nil {
		taskPush.Error = "No executable task found"
		return
	}
	err = taskService.buildTaskPush(task, &taskPush)
	return
}

func ResetAllTasks() {
	_ = global.GVA_DB.Model(&octopus.Task{}).Where("status = ?", 2).Updates(map[string]interface{}{"status": 4, "error": "服务器重启"}).Error
	return
}

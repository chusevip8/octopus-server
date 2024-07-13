package octopus

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
	"strconv"
)

type CmtTaskService struct{}

var CmtTaskServiceApp = new(CmtTaskService)

func (cmtTaskService *CmtTaskService) CreateFindCmtTask(findCmtTask *octopusReq.FindCmtTask) (err error) {
	params, err := cmtTaskService.buildFindCmtTaskParams(findCmtTask.TaskSetupId)
	if err != nil {
		return err
	}

	var taskParams octopus.TaskParams
	taskParams.TaskSetupId = findCmtTask.TaskSetupId
	taskParams.ScriptId = findCmtTask.ScriptId
	taskParams.CreatedBy = findCmtTask.CreatedBy
	taskParams.MainTaskType = findCmtTask.MainTaskType
	taskParams.SubTaskType = findCmtTask.SubTaskType
	taskParams.Params = params
	err = TaskParamsServiceApp.CreateTaskParams(&taskParams)
	if err != nil {
		return err
	}

	var task octopus.Task
	task.TaskParamsId = taskParams.ID
	task.AppName = findCmtTask.AppName
	task.DeviceId = findCmtTask.DeviceId
	task.CreatedBy = findCmtTask.CreatedBy
	task.Status = findCmtTask.Status
	task.Error = findCmtTask.Error
	err = TaskServiceApp.CreateTask(&task)
	return err
}

func (cmtTaskService *CmtTaskService) buildFindCmtTaskParams(setupId uint) (params string, err error) {
	if cmtTaskSetup, err := CmtTaskSetupServiceApp.GetCmtTaskSetup(strconv.Itoa(int(setupId))); err != nil {
		return "", err
	} else {
		postLink := cmtTaskSetup.PostLink
		keyword := cmtTaskSetup.Keyword
		params = fmt.Sprintf(`{"postLink": "%s", "keyword": "%s"}`, postLink, keyword)
		return params, nil
	}
}

func (cmtTaskService *CmtTaskService) GetTaskByDeviceId(taskSetupId string, deviceId string) (task octopus.Task, err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.task_setup_id = ? AND oct_task.device_id = ?", taskSetupId, deviceId).First(&task).Error
	return
}

func (cmtTaskService *CmtTaskService) UpdateFindCmtTaskParams(cmtTaskSetup octopus.CmtTaskSetup) (err error) {
	params := fmt.Sprintf(`{"postLink": "%s", "keyword": "%s"}`, cmtTaskSetup.PostLink, cmtTaskSetup.Keyword)
	err = global.GVA_DB.Model(&octopus.TaskParams{}).
		Where("task_setup_id = ?", cmtTaskSetup.ID).
		Where("main_task_type = ?", "cmt").
		Where("sub_task_type = ?", "findCmt").
		Update("params", params).Error
	return
}

func (cmtTaskService *CmtTaskService) DeleteCmtTask(id string, userId uint) (err error) {
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

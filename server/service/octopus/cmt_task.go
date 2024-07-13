package octopus

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"strconv"
)

type CmtTaskService struct{}

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

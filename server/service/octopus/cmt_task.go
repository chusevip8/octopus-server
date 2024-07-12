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
	params, err := cmtTaskService.buildFindCmtTaskParams(findCmtTask.SetupId)
	if err != nil {
		return err
	}

	var taskParams octopus.TaskParams
	taskParams.TaskSetupId = findCmtTask.SetupId
	taskParams.ScriptId = findCmtTask.ScriptId
	taskParams.CreatedBy = findCmtTask.CreatedBy
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
	task.Type = findCmtTask.Type
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
	db := global.GVA_DB.Model(&octopus.Task{}).Preload("TaskParams", "task_setup_id=?", taskSetupId)
	err = db.Where("device_id = ?", deviceId).First(&task).Error
	return
}

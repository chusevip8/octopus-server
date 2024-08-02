package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"strconv"
)

type GenericTaskService struct{}

func (genericTaskService *GenericTaskService) CreateGenericTask(genericTask *octopusReq.GenericTask) (err error) {

	var genericTaskSetup octopus.GenericTaskSetup

	genericTaskSetup, err = GenericTaskSetupServiceApp.GetGenericTaskSetup(strconv.Itoa(int(genericTask.TaskSetupId)))
	if err != nil {
		return err
	}
	var taskParams octopus.TaskParams
	taskParams.TaskSetupId = genericTaskSetup.ID
	taskParams.CreatedBy = genericTaskSetup.CreatedBy
	taskParams.MainTaskType = genericTask.MainTaskType
	taskParams.SubTaskType = genericTask.SubTaskType
	taskParams.Params = genericTaskSetup.Params
	taskParams.ScriptId = genericTaskSetup.ScriptId
	err = TaskParamsServiceApp.CreateTaskParams(&taskParams)
	if err != nil {
		return err
	}
	var task octopus.Task
	task.TaskParamsId = taskParams.ID
	task.AppName = genericTaskSetup.AppName
	task.DeviceId = genericTask.DeviceId
	task.CreatedBy = genericTaskSetup.CreatedBy
	task.Status = 1
	task.Error = ""
	err = TaskServiceApp.CreateTask(&task)
	return err
}

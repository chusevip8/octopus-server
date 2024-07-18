package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"strconv"
)

type IntervalTaskService struct {
}

func (intervalTaskService *IntervalTaskService) CreateIntervalTask(intervalTask *octopusReq.IntervalTask) (err error) {

	var intervalTaskSetup octopus.IntervalTaskSetup

	intervalTaskSetup, err = IntervalTaskSetupServiceApp.GetIntervalTaskSetup(strconv.Itoa(int(intervalTask.TaskSetupId)))
	if err != nil {
		return err
	}
	var taskParams octopus.TaskParams
	taskParams.TaskSetupId = intervalTaskSetup.ID
	taskParams.CreatedBy = intervalTaskSetup.CreatedBy
	taskParams.MainTaskType = intervalTask.MainTaskType
	taskParams.SubTaskType = intervalTaskSetup.SubTaskType
	taskParams.Params = intervalTaskSetup.Params
	taskParams.ScriptId = intervalTaskSetup.ScriptId
	err = TaskParamsServiceApp.CreateTaskParams(&taskParams)
	if err != nil {
		return err
	}
	var task octopus.Task
	task.TaskParamsId = taskParams.ID
	task.AppName = intervalTaskSetup.AppName
	task.DeviceId = intervalTask.DeviceId
	task.CreatedBy = intervalTaskSetup.CreatedBy
	task.Status = 1
	task.Error = ""
	err = TaskServiceApp.CreateTask(&task)
	return err
}

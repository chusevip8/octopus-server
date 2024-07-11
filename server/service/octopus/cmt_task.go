package octopus

import (
	"fmt"
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

	var cmtTaskParams octopus.CmtTaskParams
	cmtTaskParams.SetupId = findCmtTask.SetupId
	cmtTaskParams.ScriptId = findCmtTask.ScriptId
	cmtTaskParams.CreatedBy = findCmtTask.CreatedBy
	cmtTaskParams.Params = params
	err = CmtTaskParamsServiceApp.CreateCmtTaskParams(&cmtTaskParams)
	if err != nil {
		return err
	}

	var task octopus.Task
	task.TaskParamsId = cmtTaskParams.ID
	task.AppName = findCmtTask.AppName
	task.DeviceId = findCmtTask.DeviceId
	task.CreatedBy = findCmtTask.CreatedBy
	task.Type = findCmtTask.Type
	task.Status = findCmtTask.Type
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

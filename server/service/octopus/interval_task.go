package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
	"strconv"
	"time"
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
func (intervalTaskService *IntervalTaskService) DeleteIntervalTask(id string, userId uint) (err error) {
	task, err := TaskServiceApp.GetTask(id)
	if err != nil {
		return err
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		if err := tx.Model(&octopus.Task{}).Where("id = ?", task.ID).Updates(map[string]interface{}{
			"deleted_by": userId,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Model(&octopus.TaskParams{}).Where("id = ?", task.TaskParamsId).Updates(map[string]interface{}{
			"deleted_by": userId,
			"deleted_at": now,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
func (intervalTaskService *IntervalTaskService) StopIntervalTask(taskId string) (err error) {
	task, err := TaskServiceApp.GetTask(taskId)
	if err != nil {
		return err
	}
	StopTask <- &task
	return nil
}

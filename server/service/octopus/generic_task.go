package octopus

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"strconv"
)

type GenericTaskService struct{}

func (genericTaskService *GenericTaskService) CreateGenericTask(genericTask *octopusReq.GenericTask) (err error) {
	var genericTaskSetup octopus.GenericTaskSetup

	// 获取任务设置
	genericTaskSetup, err = GenericTaskSetupServiceApp.GetGenericTaskSetup(strconv.Itoa(int(genericTask.TaskSetupId)))
	if err != nil {
		return err
	}

	// 任务参数通用部分填充
	fillTaskParams := func() (octopus.TaskParams, error) {
		taskParams := octopus.TaskParams{
			TaskSetupId:  genericTaskSetup.ID,
			CreatedBy:    genericTaskSetup.CreatedBy,
			MainTaskType: genericTask.MainTaskType,
			SubTaskType:  genericTask.SubTaskType,
			Params:       "",
			ScriptId:     genericTaskSetup.ScriptId,
		}
		err := TaskParamsServiceApp.CreateTaskParams(&taskParams)
		return taskParams, err
	}

	// 创建任务
	createTask := func(deviceId uint, taskParamsId uint) error {
		task := octopus.Task{
			TaskParamsId: taskParamsId,
			AppName:      genericTaskSetup.AppName,
			DeviceId:     deviceId,
			CreatedBy:    genericTaskSetup.CreatedBy,
			Status:       1,
			Error:        "",
		}
		return TaskServiceApp.CreateTaskWithoutNotify(&task)
	}

	if genericTask.Batch {
		// 批量任务创建
		devices, err := DeviceServiceApp.GetReadyDeviceListByUserId(genericTask.CreatedBy, genericTask.DeviceGroup)
		if err != nil {
			return err
		}

		for _, device := range devices {
			taskParams, err := fillTaskParams()
			if err != nil {
				return err
			}
			if err = createTask(device.ID, taskParams.ID); err != nil {
				return err
			}
		}
	} else {
		// 单个任务创建
		taskParams, err := fillTaskParams()
		if err != nil {
			return err
		}
		err = createTask(genericTask.DeviceId, taskParams.ID)
	}

	return err
}

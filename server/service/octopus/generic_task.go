package octopus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusReq "github.com/flipped-aurora/gin-vue-admin/server/model/octopus/request"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"strconv"
)

type GenericTaskService struct{}

func (genericTaskService *GenericTaskService) fillTaskParams(genericTaskSetup octopus.GenericTaskSetup, genericTask octopusReq.GenericTask, params string) (octopus.TaskParams, error) {
	taskParams := octopus.TaskParams{
		TaskSetupId:  genericTaskSetup.ID,
		CreatedBy:    genericTaskSetup.CreatedBy,
		MainTaskType: genericTask.MainTaskType,
		SubTaskType:  genericTask.SubTaskType,
		Params:       params,
		ScriptId:     genericTaskSetup.ScriptId,
	}
	err := TaskParamsServiceApp.CreateTaskParams(&taskParams)
	return taskParams, err
}

func (genericTaskService *GenericTaskService) createTask(genericTaskSetup octopus.GenericTaskSetup, deviceId uint, taskParamsId uint) error {
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

func (genericTaskService *GenericTaskService) CreateGenericTask(genericTask octopusReq.GenericTask) (err error) {

	// 获取任务设置
	genericTaskSetup, err := GenericTaskSetupServiceApp.GetGenericTaskSetup(strconv.Itoa(int(genericTask.TaskSetupId)))
	if err != nil {
		return err
	}
	script, err := ScriptServiceApp.GetScript(strconv.Itoa(int(genericTaskSetup.ScriptId)))
	if err != nil {
		return err
	}
	scriptKeywords := genericTaskService.findScriptKeywords(script.Content)

	if genericTask.Batch {
		// 批量任务创建
		devices, err := DeviceServiceApp.GetReadyDeviceListByUserId(genericTask.CreatedBy, genericTask.DeviceGroup)
		if err != nil {
			return err
		}

		for _, device := range devices {
			bindData, err := taskBindDataServiceApp.GetNewBindData(strconv.Itoa(int(genericTaskSetup.ID)), genericTask.MainTaskType)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			} else if err != nil {
				return err
			}
			params, err := genericTaskService.buildTaskParams(scriptKeywords, bindData)
			if err != nil {
				return err
			}
			taskParams, err := genericTaskService.fillTaskParams(genericTaskSetup, genericTask, params)
			if err != nil {
				return err
			}
			if err = genericTaskService.createTask(genericTaskSetup, device.ID, taskParams.ID); err != nil {
				return err
			}
		}
	} else {
		// 单个任务创建
		bindData, err := taskBindDataServiceApp.GetNewBindData(strconv.Itoa(int(genericTaskSetup.ID)), genericTask.MainTaskType)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else if err != nil {
			return err
		}
		params, err := genericTaskService.buildTaskParams(scriptKeywords, bindData)
		if err != nil {
			return err
		}
		taskParams, err := genericTaskService.fillTaskParams(genericTaskSetup, genericTask, params)
		if err != nil {
			return err
		}
		err = genericTaskService.createTask(genericTaskSetup, genericTask.DeviceId, taskParams.ID)
	}

	return err
}
func (genericTaskService *GenericTaskService) buildTaskParams(scriptKeywords []string, bindData octopus.TaskBindData) (params string, err error) {
	if len(scriptKeywords) == 0 {
		return "", err
	}
	resultMap := make(map[string]string)

	val := reflect.ValueOf(bindData)

	fieldNames := []string{
		"Item1", "Item2", "Item3", "Item4", "Item5",
		"Item6", "Item7", "Item8", "Item9", "Item10",
	}
	for i, keyword := range scriptKeywords {
		if i >= len(fieldNames) {
			return "", fmt.Errorf("scriptKeywords count exceeds available fields in TaskBindData")
		}

		fieldName := fieldNames[i]

		fieldVal := val.FieldByName(fieldName).String()

		resultMap[keyword] = fieldVal
	}
	jsonData, err := json.Marshal(resultMap)
	if err != nil {
		return "", fmt.Errorf("failed to marshal resultMap to JSON: %v", err)
	}
	return string(jsonData), nil
}

func (genericTaskService *GenericTaskService) findScriptKeywords(scriptContent string) []string {
	// 正则表达式，用于匹配 ${} 并捕获 {} 内的内容
	re := regexp.MustCompile(`\$\{([^}]*)}`)
	matches := re.FindAllStringSubmatch(scriptContent, -1)

	var keywords []string
	for _, match := range matches {
		if len(match) > 1 {
			keywords = append(keywords, match[1]) // match[1] 是捕获组中的内容
		}
	}
	return keywords
}

func (genericTaskService *GenericTaskService) BindTaskData(bindTaskData octopusReq.BindTaskData) (err error) {
	tasks, err := TaskServiceApp.GetTasksByTaskSetupId(bindTaskData.TaskSetupId, bindTaskData.MainTaskType, bindTaskData.SubTaskType)
	if err != nil {
		return err
	} else if len(tasks) == 0 {
		return fmt.Errorf("dvice not found task with id %s", bindTaskData.TaskSetupId)
	}
	// 获取任务设置
	genericTaskSetup, err := GenericTaskSetupServiceApp.GetGenericTaskSetup(bindTaskData.TaskSetupId)
	if err != nil {
		return err
	}
	script, err := ScriptServiceApp.GetScript(strconv.Itoa(int(genericTaskSetup.ScriptId)))
	if err != nil {
		return err
	}
	scriptKeywords := genericTaskService.findScriptKeywords(script.Content)
	taskIndex := 0
	for {
		bindData, err := taskBindDataServiceApp.GetNewBindData(strconv.Itoa(int(genericTaskSetup.ID)), bindTaskData.MainTaskType)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		deviceId := tasks[taskIndex].DeviceId
		params, err := genericTaskService.buildTaskParams(scriptKeywords, bindData)
		if err != nil {
			return err
		}
		genericTask := octopusReq.GenericTask{MainTaskType: bindTaskData.MainTaskType, SubTaskType: bindTaskData.SubTaskType}
		taskParams, err := genericTaskService.fillTaskParams(genericTaskSetup, genericTask, params)
		if err != nil {
			return err
		}
		err = genericTaskService.createTask(genericTaskSetup, deviceId, taskParams.ID)
		if err != nil {
			return err
		}
		taskIndex = (taskIndex + 1) % len(tasks)
	}
}

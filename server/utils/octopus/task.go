package octopus

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"strconv"
	"strings"
)

func DeviceIsReady(deviceId string) bool {
	var task octopus.Task
	err := global.GVA_DB.Model(&octopus.Task{}).
		Where("device_id = ?", deviceId).
		Where("status = ?", 2).
		First(&task).Error
	if err != nil {
		var device *octopus.Device
		err = global.GVA_DB.Model(&octopus.Device{}).
			Where("id = ?", deviceId).
			Where("status = ?", 1).
			First(&device).Error
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func findPushTask(deviceId string) (task octopus.Task, err error) {
	err = global.GVA_DB.Model(&octopus.Task{}).
		Joins("LEFT JOIN oct_task_params ON oct_task_params.id = oct_task.task_params_id").
		Where("oct_task_params.main_task_type != ? AND oct_task.device_id = ?", "interval", deviceId).
		First(&task).Error
	return
}

func buildTaskPush(task octopus.Task, taskPush *protocol.TaskPush) (err error) {

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

func PushTask(deviceId string) (taskPush protocol.TaskPush, err error) {
	task, err := findPushTask(deviceId)
	if err != nil {
		taskPush.Error = "No executable task found"
		return
	}
	err = buildTaskPush(task, &taskPush)
	return
}

func UpdateAllTasks() {
	_ = global.GVA_DB.Model(&octopus.Task{}).Where("status = ?", 2).Updates(map[string]interface{}{"status": 4, "error": "服务器重启"}).Error
	return
}

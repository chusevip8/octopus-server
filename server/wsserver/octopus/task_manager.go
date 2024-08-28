package octopus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusService "github.com/flipped-aurora/gin-vue-admin/server/service/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

func tryPushTask(task *octopus.Task) {
	global.GVA_LOG.Info("New Task created", zap.String("Id", strconv.Itoa(int(task.ID))))
	client, ok := socket.GetClient(task.DeviceId)
	if ok {
		client.ClientLock.Lock()
		defer client.ClientLock.Unlock()
		deviceId := strconv.Itoa(int(task.DeviceId))
		ready := service.DeviceService.DeviceIsFree(deviceId)
		if ready {
			pushTaskId, data, err := PushTaskMessage(deviceId)
			if err != nil {
				global.GVA_LOG.Error("Try push task message", zap.String("error", err.Error()))
				_ = service.TaskService.UpdateTaskStatusToFail(pushTaskId, "Try push task message"+err.Error())
			} else {
				err = service.TaskService.UpdateTaskStatusToRun(pushTaskId)
				if err != nil {
					global.GVA_LOG.Error("Try push task update task status", zap.String("error", err.Error()))
				} else {
					client.SendMessage(data)
				}
			}
		}
	}
}

func ResetAllTasks() {
	service.TaskService.UpdateTaskStatusRunToFail()
}

func PushTaskMessage(deviceId string) (pushTaskId string, data []byte, err error) {
	var taskPush protocol.TaskPush
	task, err := service.TaskService.FindPushTask(deviceId)
	if err != nil {
		return
	}
	err = buildTaskPush(task, &taskPush)
	if err != nil {
		return
	}
	pushTaskId = taskPush.TaskId
	message := map[string]interface{}{"code": protocol.CodeTaskPush, "data": taskPush}
	data, err = json.Marshal(message)
	return
}

func buildTaskPush(task octopus.Task, taskPush *protocol.TaskPush) (err error) {
	script, err := service.ScriptService.GetScript(strconv.Itoa(int(task.TaskParams.ScriptId)))
	if err != nil {
		taskPush.Error = "Can't find task script"
		global.GVA_LOG.Error("Build task push,Can't find task script", zap.String("error", err.Error()))
		return
	}
	scriptContent := script.Content
	if task.TaskParams.Params != "" {
		var params map[string]string
		err = json.Unmarshal([]byte(task.TaskParams.Params), &params)
		if err != nil {
			taskPush.Error = "Task params json unmarshal error"
			global.GVA_LOG.Error("Build task push,Task params json unmarshal error", zap.String("error", err.Error()))
			return
		}
		for key, value := range params {
			placeholder := fmt.Sprintf("${%s}", key)
			scriptContent = strings.ReplaceAll(scriptContent, placeholder, value)
		}
	}
	var compactBuffer bytes.Buffer
	if err = json.Compact(&compactBuffer, []byte(scriptContent)); err != nil {
		taskPush.Error = "Task script compact error"
		global.GVA_LOG.Error("Build task push,Task script compact error", zap.String("error", err.Error()))
		return
	}
	taskPush.TaskId = strconv.Itoa(int(task.ID))
	taskPush.Script = compactBuffer.String()
	taskPush.Error = ""
	return
}

func tryStopTask(task *octopus.Task) {
	global.GVA_LOG.Info("Stop Task ", zap.String("Id", strconv.Itoa(int(task.ID))))
	client, ok := socket.GetClient(task.DeviceId)
	if ok {
		client.ClientLock.Lock()
		defer client.ClientLock.Unlock()
		global.GVA_LOG.Info("Trying to stop task", zap.String("Task status", strconv.Itoa(int(task.Status))))
		if task.Status == 1 {
			_ = service.TaskService.UpdateTaskStatusToFail(strconv.Itoa(int(task.ID)), "服务器终止任务")
		} else if task.Status == 2 {
			taskStopPush := protocol.TaskStopPush{TaskId: strconv.Itoa(int(task.ID)), Error: "服务器终止任务"}
			message := map[string]interface{}{"code": protocol.CodeTaskStopPush, "data": taskStopPush}
			data, err := json.Marshal(message)
			if err != nil {
				global.GVA_LOG.Error("Try push stop task message ", zap.String("task", strconv.Itoa(int(task.ID))), zap.String("error", err.Error()))
			} else {
				client.SendMessage(data)
			}
		}
	} else {
		if task.Status == 1 || task.Status == 2 {
			_ = service.TaskService.UpdateTaskStatusToFail(strconv.Itoa(int(task.ID)), "服务器终止任务")
		}
	}
}
func IntervalTaskMessage(task octopus.Task) (pushTaskId string, data []byte, err error) {

	var taskPush protocol.TaskPush
	if err = buildTaskPush(task, &taskPush); err != nil {
		return "", nil, err
	}

	pushTaskId = taskPush.TaskId
	message := map[string]interface{}{"code": protocol.CodeTaskPush, "data": taskPush}
	if data, err = json.Marshal(message); err != nil {
		return "", nil, err
	}

	return pushTaskId, data, nil
}
func TryPushIntervalTask(task octopus.Task) {
	client, exists := socket.GetClient(task.DeviceId)
	if !exists {
		global.GVA_LOG.Info("device not found", zap.String("device", strconv.Itoa(int(task.DeviceId))))
		return
	}

	client.ClientLock.Lock()
	defer client.ClientLock.Unlock()

	deviceId := strconv.Itoa(int(task.DeviceId))
	if !service.DeviceService.DeviceIsFree(deviceId) {
		global.GVA_LOG.Info("device is not ready", zap.String("device", strconv.Itoa(int(task.DeviceId))))
		return
	}

	pushTaskId, data, err := IntervalTaskMessage(task)
	if err != nil {
		global.GVA_LOG.Error("Try push interval task message", zap.String("error", err.Error()))
	} else {
		err = service.TaskService.UpdateTaskStatusToRun(pushTaskId)
		if err != nil {
			global.GVA_LOG.Error("Try push interval task update task status", zap.String("error", err.Error()))
		} else {
			global.GVA_LOG.Info("interval task is running", zap.Any("task", task))
			client.SendMessage(data)
		}
	}
}

func MonitorTask() {
	go func() {
		for {
			select {
			case newTask := <-octopusService.NewTask:
				go tryPushTask(newTask)
			case stopTask := <-octopusService.StopTask:
				go tryStopTask(stopTask)
			}
		}
	}()
}

package octopus

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/octopus"
	octopusService "github.com/flipped-aurora/gin-vue-admin/server/service/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"strconv"
	"strings"
)

func tryPushTask(task *octopus.Task) {
	fmt.Printf("New Task created:(ID: %d)\n", task.ID)
	client, ok := socket.GetClient(task.DeviceId)
	if ok {
		client.ClientLock.Lock()
		defer client.ClientLock.Unlock()
		deviceId := strconv.Itoa(int(task.DeviceId))
		ready := service.DeviceService.DeviceIsFree(deviceId)
		if ready {
			data, err := PushTaskMessage(deviceId)
			if err != nil {
				fmt.Println("Try push task message", err)
			} else {
				err = service.TaskService.UpdateTaskStatusToRun(deviceId)
				if err != nil {
					fmt.Println("Try push task update task status", err)
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

func PushTaskMessage(deviceId string) (data []byte, err error) {
	var taskPush protocol.TaskPush
	task, err := service.TaskService.FindPushTask(deviceId)
	if err != nil {
		return
	}
	err = buildTaskPush(task, &taskPush)
	if err != nil {
		return
	}
	message := map[string]interface{}{"code": protocol.CodeTaskPush, "data": taskPush}
	data, err = json.Marshal(message)
	return
}

func buildTaskPush(task octopus.Task, taskPush *protocol.TaskPush) (err error) {
	script, err := service.ScriptService.GetScript(strconv.Itoa(int(task.TaskParams.ScriptId)))
	if err != nil {
		taskPush.Error = "Can't find task script"
		return
	}
	scriptContent := script.Content
	if task.TaskParams.Params != "" {
		var params map[string]string
		err = json.Unmarshal([]byte(task.TaskParams.Params), &params)
		if err != nil {
			taskPush.Error = "Task params json unmarshal error"
			return
		}
		for key, value := range params {
			placeholder := fmt.Sprintf("${%s}", key)
			scriptContent = strings.ReplaceAll(scriptContent, placeholder, value)
		}
	}
	taskPush.TaskId = strconv.Itoa(int(task.ID))
	taskPush.Script = scriptContent
	taskPush.Error = ""
	return
}

func WaitNewTask() {
	go func() {
		for {
			select {
			case task := <-octopusService.NewTask:
				go tryPushTask(task)
			}
		}
	}()
}

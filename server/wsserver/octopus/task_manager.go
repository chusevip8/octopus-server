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
	deviceId := strconv.Itoa(int(task.DeviceId))
	ready := service.DeviceService.DeviceIsFree(deviceId)
	if ready {
		client, ok := socket.GetClient(task.DeviceId)
		if ok {
			taskPush, err := PushTask(deviceId)
			if err != nil {
				fmt.Println("After task create", err)
			} else {
				message := map[string]interface{}{"code": protocol.CodeTaskPush, "data": taskPush}
				data, err := json.Marshal(message)
				if err != nil {
					fmt.Println("After task create marshal", err)
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

func PushTask(deviceId string) (taskPush protocol.TaskPush, err error) {
	task, err := service.TaskService.FindPushTask(deviceId)
	if err != nil {
		taskPush.Error = "No executable task found"
		return
	}
	err = buildTaskPush(task, &taskPush)
	return
}

func buildTaskPush(task octopus.Task, taskPush *protocol.TaskPush) (err error) {
	script, err := service.ScriptService.GetScript(strconv.Itoa(int(task.TaskParams.ScriptId)))
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

func WaitNewTask() {
	for {
		select {
		case task := <-octopusService.NewTask:
			go tryPushTask(task)
		}
	}

}

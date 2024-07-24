package socket

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"strconv"
)

func TaskFinishHandler(client *Client, data []byte) {
	taskFinish := &protocol.TaskFinish{}
	if err := json.Unmarshal(data, taskFinish); err != nil {
		fmt.Println("TaskFinishHandler json Unmarshal", err)
		return
	}
	if taskFinish.Error == "" {
		_ = taskService.UpdateTaskStatusToFinish(taskFinish.TaskId)
	} else {
		_ = taskService.UpdateTaskStatusToFail(taskFinish.TaskId, taskFinish.Error)
	}

	deviceId := strconv.Itoa(int(client.Id))
	taskPush, _ := taskService.PushTask(deviceId)
	message := map[string]interface{}{"code": protocol.CodeTaskPush, "data": taskPush}
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("RequestTaskHandler", err)
	} else {
		client.SendMessage(data)
	}
}

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
)

func TaskFinishHandler(client *socket.Client, data []byte) {
	taskFinish := &protocol.TaskFinish{}
	if err := json.Unmarshal(data, taskFinish); err != nil {
		fmt.Println("TaskFinishHandler json Unmarshal", err)
		return
	}
	if taskFinish.Error == "" {
		_ = service.TaskService.UpdateTaskStatusToFinish(taskFinish.TaskId)
	} else {
		_ = service.TaskService.UpdateTaskStatusToFail(taskFinish.TaskId, taskFinish.Error)
	}

	taskFinishPush := protocol.TaskFinishPush{Token: taskFinish.Token, TaskId: taskFinish.TaskId}

	message := map[string]interface{}{"code": protocol.CodeTaskFinishPush, "data": taskFinishPush}
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("TaskFinishHandler", err)
	} else {
		client.SendMessage(data)
	}
}

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
)

func TaskStartHandler(client *socket.Client, data []byte) {
	taskStart := &protocol.TaskStart{}
	if err := json.Unmarshal(data, taskStart); err != nil {
		fmt.Println("TaskStartHandler json Unmarshal", err)
		return
	}
	_ = octopus.TaskService.UpdateTaskStatusToRun(taskStart.TaskId)
}

package socket

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
)

func TaskStartHandler(client *Client, data []byte) {
	taskStart := &protocol.TaskStart{}
	if err := json.Unmarshal(data, taskStart); err != nil {
		fmt.Println("TaskStartHandler json Unmarshal", err)
		return
	}
	_ = taskService.UpdateTaskStatusToRun(taskStart.TaskId)
}

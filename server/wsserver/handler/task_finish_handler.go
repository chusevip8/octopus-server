package handler

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"go.uber.org/zap"
)

func TaskFinishHandler(client *socket.Client, data []byte) {
	taskFinish := &protocol.TaskFinish{}
	if err := json.Unmarshal(data, taskFinish); err != nil {
		global.GVA_LOG.Error("TaskFinishHandler json Unmarshal", zap.String("error", err.Error()))
		return
	}
	client.ClientLock.Lock()
	defer client.ClientLock.Unlock()
	if taskFinish.Error == "" {
		_ = service.TaskService.UpdateTaskStatusToFinish(taskFinish.TaskId)
	} else {
		_ = service.TaskService.UpdateTaskStatusToFail(taskFinish.TaskId, taskFinish.Error)
	}

	taskFinishPush := protocol.TaskFinishPush{Token: taskFinish.Token, TaskId: taskFinish.TaskId}

	message := map[string]interface{}{"code": protocol.CodeTaskFinishPush, "data": taskFinishPush}
	data, err := json.Marshal(message)
	if err != nil {
		global.GVA_LOG.Error("TaskFinishHandler json marshal", zap.String("error", err.Error()))
	} else {
		client.SendMessage(data)
	}
}

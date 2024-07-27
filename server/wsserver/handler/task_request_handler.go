package handler

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"go.uber.org/zap"
	"strconv"
)

func TaskRequestHandler(client *socket.Client, data []byte) {
	deviceId := strconv.Itoa(int(client.Id))
	client.ClientLock.Lock()
	defer client.ClientLock.Unlock()

	data, err := octopus.PushTaskMessage(deviceId)
	if err != nil {
		handleTaskPushError(client, err, "Request task handler message")
		return
	}

	if err := service.TaskService.UpdateTaskStatusToRun(deviceId); err != nil {
		handleTaskPushError(client, err, "Request task handler update task status")
		return
	}

	client.SendMessage(data)
}
func handleTaskPushError(client *socket.Client, err error, context string) {
	global.GVA_LOG.Error(context, zap.String("error", err.Error()))
	message := map[string]interface{}{
		"code": protocol.CodeTaskPush,
		"data": protocol.TaskPush{TaskId: "", Script: "", Error: err.Error()},
	}
	data, jsonErr := json.Marshal(message)
	if jsonErr == nil {
		client.SendMessage(data)
	}
}

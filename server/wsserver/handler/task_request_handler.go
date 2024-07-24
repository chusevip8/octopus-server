package handler

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"strconv"
)

func TaskRequestHandler(client *socket.Client, data []byte) {
	deviceId := strconv.Itoa(int(client.Id))
	taskPush, _ := octopus.TaskService.PushTask(deviceId)
	message := map[string]interface{}{"code": protocol.CodeTaskPush, "data": taskPush}
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("RequestTaskHandler", err)
	} else {
		client.SendMessage(data)
	}

}

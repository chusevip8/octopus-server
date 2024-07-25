package handler

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"strconv"
)

func TaskRequestHandler(client *socket.Client, data []byte) {
	deviceId := strconv.Itoa(int(client.Id))
	client.ClientLock.Lock()
	defer client.ClientLock.Unlock()
	data, err := octopus.PushTaskMessage(deviceId)
	if err != nil {
		fmt.Println("Request task handler message", err)
	} else {
		err = service.TaskService.UpdateTaskStatusToRun(deviceId)
		if err != nil {
			fmt.Println("Request task handler update task status", err)
		} else {
			client.SendMessage(data)
		}
	}
}

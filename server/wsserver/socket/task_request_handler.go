package socket

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"strconv"
)

func TaskRequestHandler(client *Client, data []byte) {
	deviceId := strconv.Itoa(int(client.Id))
	taskPush, _ := octopus.PushTask(deviceId)
	message := map[string]interface{}{"code": protocol.CodeTaskPush, "data": taskPush}
	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("RequestTaskHandler", err)
	} else {
		client.SendMessage(data)
	}

}

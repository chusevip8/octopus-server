package socket

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
)

func LoginHandler(client *Client, data []byte) {

	login := &protocol.Login{}
	if err := json.Unmarshal(data, login); err != nil {
		fmt.Println("LoginHandler json Unmarshal", err)
		return
	}

	var loginPush protocol.LoginPush

	device, err := deviceService.GetDeviceByToken(login.Token)
	if err != nil {
		client.Id = 0
		loginPush.Code = 1
		loginPush.Error = "Device not found"
	} else {
		client.Id = device.ID
		loginPush.Code = 0
		loginPush.Error = "Device login"
	}
	message, err := json.Marshal(loginPush)
	if err != nil {
		fmt.Println("LoginHandler json Marshal", err)
		return
	}
	client.SendMessage(message)
	if loginPush.Code != 0 {
		client.Close()
	} else {
		client.Hub.Login <- client
	}
}

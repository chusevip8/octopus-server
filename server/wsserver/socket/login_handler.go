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
		loginPush.Token = login.Token
		loginPush.Error = "Device not found"
	} else {
		client.Id = device.ID
		loginPush.Token = login.Token
		loginPush.Error = ""
	}
	message, err := json.Marshal(loginPush)
	if err != nil {
		fmt.Println("LoginHandler json Marshal", err)
		return
	}
	client.SendMessage(message)
	if loginPush.Error != "" {
		client.Close()
	} else {
		client.Hub.Login <- client
	}
}

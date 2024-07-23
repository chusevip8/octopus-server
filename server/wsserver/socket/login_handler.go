package socket

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
)

const Login = 2

func LoginHandler(client *Client, data []byte) {

	loginReq := &protocol.LoginReq{}
	if err := json.Unmarshal(data, loginReq); err != nil {
		fmt.Println("LoginHandler json Unmarshal", err)
		return
	}

	var loginRes protocol.LoginRes

	device, err := deviceService.GetDeviceByToken(loginReq.Token)
	if err != nil {
		client.Id = 0
		loginRes.Code = 1
		loginRes.Error = "Device not found"
	} else {
		client.Id = device.ID
		loginRes.Code = 0
		loginRes.Error = "Device login"
	}
	message, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("LoginHandler json Marshal", err)
		return
	}
	client.SendMessage(message)
	if loginRes.Code != 0 {
		client.Close()
	} else {
		client.Hub.Login <- client
	}
}

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
)

func LoginHandler(client *socket.Client, data []byte) {

	var loginPush protocol.LoginPush
	login := &protocol.Login{}
	if err := json.Unmarshal(data, login); err != nil {
		fmt.Println("LoginHandler json Unmarshal", err)
		client.Id = 0
		loginPush.Token = login.Token
		loginPush.Error = err.Error()
	} else {
		device, err := service.DeviceService.GetDeviceByToken(login.Token)
		if err != nil {
			fmt.Println("LoginHandler device not found", err)
			client.Id = 0
			loginPush.Token = login.Token
			loginPush.Error = err.Error()
		} else {
			client.Id = device.ID
			loginPush.Token = login.Token
			loginPush.Error = ""
		}
	}

	message, err := json.Marshal(protocol.Message{Code: protocol.CodeLoginPush, Data: loginPush})
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

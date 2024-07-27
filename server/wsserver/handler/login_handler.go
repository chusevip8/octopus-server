package handler

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"go.uber.org/zap"
)

func LoginHandler(client *socket.Client, data []byte) {

	var loginPush protocol.LoginPush
	login := &protocol.Login{}
	if err := json.Unmarshal(data, login); err != nil {
		global.GVA_LOG.Info("Login Handler json Unmarshal", zap.String("error", err.Error()))
		client.Id = 0
		loginPush.Token = login.Token
		loginPush.Error = err.Error()
	} else {
		device, err := service.DeviceService.GetDeviceByToken(login.Token)
		if err != nil {
			global.GVA_LOG.Info("Login device not found", zap.String("token", login.Token))
			client.Id = 0
			loginPush.Token = login.Token
			loginPush.Error = err.Error()
		} else {
			client.Id = device.ID
			client.UserName = device.Username
			client.Number = device.Number
			loginPush.Token = login.Token
			loginPush.Error = ""
		}
	}

	message, err := json.Marshal(protocol.Message{Code: protocol.CodeLoginPush, Data: loginPush})
	if err != nil {
		global.GVA_LOG.Info("Login Handler json Marshal", zap.String("error", err.Error()))
		return
	}
	client.SendMessage(message)
	if loginPush.Error != "" {
		client.Close()
	} else {
		client.Hub.Login <- client
	}
}

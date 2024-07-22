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

	_, err := deviceService.GetDeviceByToken(loginReq.Token)
	if err != nil {
		loginRes.Code = -1
		loginRes.Error = "未找到该设备"
	} else {
		loginRes.Code = 0
		loginRes.Error = ""
	}
	message, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("LoginHandler json Marshal", err)
		return
	}
	client.SendMessage(message)
	if loginRes.Code != 0 {
		client.SendMessage([]byte(protocol.CloseSignal))
	}
}

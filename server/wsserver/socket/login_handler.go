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

	//service.DeviceService
	var loginRes protocol.LoginRes
	loginRes.Code = 1
	loginRes.Error = ""
	message, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("LoginHandler json Marshal", err)
		return
	}
	client.SendMessage(message)
	client.SendMessage([]byte(protocol.CloseSignal))
}

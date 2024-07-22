package socket

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"time"
)

type Hub struct {
	Connect    chan *Client
	Disconnect chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Connect:    make(chan *Client),
		Disconnect: make(chan *Client),
	}
}

func (hub *Hub) Run(clientManager *ClientManager) {
	for {
		select {
		case client := <-hub.Connect:
			go hub.checkClientLogin(client)
		case client := <-hub.Disconnect:
			go hub.removeClient(client, clientManager)
		}
	}
}

func (hub *Hub) checkClientLogin(client *Client) {
	select {
	case <-time.After(10 * time.Second):
		if client != nil && client.Id == 0 {
			client.SendMessage([]byte(protocol.CloseSignal))
		}
	}
}

func (hub *Hub) removeClient(client *Client, clientManager *ClientManager) {
	clientManager.RemoveClient(client)
	close(client.Send)
}

func (hub *Hub) ReceiveMessage(client *Client, data []byte) {
	fmt.Println("Receive Message", client.Addr, string(data))
	message := &protocol.Message{}
	if err := json.Unmarshal(data, message); err != nil {
		fmt.Println("Receive Message json Unmarshal", err)
		return
	}
	msgContent, err := json.Marshal(message.Data)
	if err != nil {
		fmt.Println("Receive Message json Marshal", err)
		return
	}
	if handle, ok := GetHandler(message.Code); ok {
		handle(client, msgContent)
	} else {
		fmt.Println("Handler not found", client.Addr, "Code", message.Code)
	}
}
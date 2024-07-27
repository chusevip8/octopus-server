package socket

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/service"
	"go.uber.org/zap"
	"time"
)

type MessageReceiver interface {
	HandleMessage(client *Client, data []byte)
}

type Hub struct {
	Connect        chan *Client
	Disconnect     chan *Client
	Login          chan *Client
	MessageHandler MessageReceiver
}

func NewHub(handler MessageReceiver) *Hub {
	return &Hub{
		Connect:        make(chan *Client),
		Disconnect:     make(chan *Client),
		Login:          make(chan *Client),
		MessageHandler: handler,
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Connect:
			go hub.checkClientLogin(client)
		case client := <-hub.Disconnect:
			go hub.disconnect(client)
		case client := <-hub.Login:
			go hub.login(client)
		}
	}
}

func (hub *Hub) disconnect(client *Client) {
	fmt.Println("Disconnected:", client.Conn.RemoteAddr().String())
	_ = service.DeviceService.UpdateDeviceStatusById(client.Id, 2)
	_ = service.TaskService.UpdateTaskStatusRunToFailByDeviceId(client.Id, "设备离线")
	RemoveClient(client)
	close(client.Send)

}
func (hub *Hub) login(client *Client) {
	fmt.Println("Logged in:", client.Conn.RemoteAddr().String())
	_ = service.DeviceService.UpdateDeviceStatusById(client.Id, 1)
	AddClient(client)
}

func (hub *Hub) checkClientLogin(client *Client) {
	global.GVA_LOG.Info("Connected", zap.String("address", client.Conn.RemoteAddr().String()))
	select {
	case <-time.After(10 * time.Second):
		if client != nil && client.Id == 0 {
			client.Close()
		}
	}
}

func (hub *Hub) ReceiveMessage(client *Client, data []byte) {
	fmt.Println("Receive Message", client.Addr, string(data))
	if hub.MessageHandler != nil {
		hub.MessageHandler.HandleMessage(client, data)
	}
}

package socket

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"log"
	"sync"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pingWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pongPeriod = (pingWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	Hub        *Hub
	Id         uint
	UserName   string
	Number     string
	Addr       string
	IsClose    bool
	Conn       *websocket.Conn
	Send       chan []byte
	ClientLock sync.Mutex
}

func NewClient(hub *Hub, addr string, conn *websocket.Conn) (client *Client) {
	client = &Client{
		Hub:      hub,
		Id:       0,
		UserName: "",
		Number:   "",
		Addr:     addr,
		IsClose:  false,
		Conn:     conn,
		Send:     make(chan []byte, 256),
	}
	return
}
func (client *Client) Read() {
	defer func() {
		_ = client.Conn.Close()
		client.IsClose = true
		client.Hub.Disconnect <- client
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	_ = client.Conn.SetReadDeadline(time.Now().Add(pingWait))
	client.Conn.SetPingHandler(func(string) error {
		fmt.Println("Receive ping message", client.Conn.RemoteAddr())
		_ = client.Conn.SetReadDeadline(time.Now().Add(pingWait))
		_ = client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
		err := client.Conn.WriteMessage(websocket.PongMessage, nil)
		return err
	})
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		rawData, err := octopus.DataIn(message)
		if err != nil {
			fmt.Println("Receive data", err)
		} else {
			client.Hub.ReceiveMessage(client, rawData)
		}
	}
}

func (client *Client) Write() {
	defer func() {
		_ = client.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				_ = client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if string(message) == protocol.CloseSignal {
				_ = client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			_ = client.Conn.WriteMessage(websocket.BinaryMessage, message)
		}
	}
}
func (client *Client) SendMessage(data []byte) {
	outData, err := octopus.DataOut(data)
	if err != nil {
		global.GVA_LOG.Error("Send message", zap.String("message", string(data)),
			zap.String("clientUsername", client.UserName),
			zap.String("clientNumber", client.Number),
			zap.String("clientAddr", client.Addr),
			zap.String("error", err.Error()))
		return
	}
	if client.IsClose {
		return
	}
	client.Send <- outData
	global.GVA_LOG.Info("Send message", zap.String("message", string(data)),
		zap.String("clientUsername", client.UserName),
		zap.String("clientNumber", client.Number),
		zap.String("clientAddr", client.Addr))

}
func (client *Client) Close() {
	if client.IsClose {
		return
	}
	client.Send <- []byte(protocol.CloseSignal)
	global.GVA_LOG.Info("Send Close signal",
		zap.String("clientUsername", client.UserName),
		zap.String("clientNumber", client.Number),
		zap.String("clientAddr", client.Addr))
}

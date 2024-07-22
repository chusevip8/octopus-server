package socket

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/octopus"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/gorilla/websocket"
	"log"
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
	Hub  *Hub
	Id   uint
	Addr string
	Conn *websocket.Conn
	Send chan []byte
}

func NewClient(hub *Hub, addr string, conn *websocket.Conn) (client *Client) {
	client = &Client{
		Hub:  hub,
		Id:   0,
		Addr: addr,
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	return
}
func (client *Client) Read() {
	defer func() {
		client.Hub.Disconnect <- client
		_ = client.Conn.Close()
	}()
	client.Conn.SetReadLimit(maxMessageSize)
	_ = client.Conn.SetReadDeadline(time.Now().Add(pingWait))
	client.Conn.SetPingHandler(func(string) error {
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
			_ = client.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
func (client *Client) SendMessage(data []byte) {
	if client == nil {
		return
	}
	outData, err := octopus.DataOut(data)
	if err != nil {
		log.Printf("send data error: %v", err)
		return
	}
	client.Send <- outData
}
func (client *Client) Close() {
	if client == nil {
		return
	}
	client.Send <- []byte(protocol.CloseSignal)
}

package wsserver

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/handler"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println("http upgrade", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func run() {
	addr := "127.0.0.1:5555"
	handlerManager := handler.NewHandlerManager()
	handlerManager.RegisterAllHandlers()
	hub := socket.NewHub(handlerManager)
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	fmt.Println("ws server running on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Start() {
	go run()
}

func serveWs(hub *socket.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Failed to upgrade to WebSocket", err)
		return
	}
	fmt.Println("webSocket connected:", conn.RemoteAddr().String())
	client := socket.NewClient(hub, conn.RemoteAddr().String(), conn)
	go client.Read()
	go client.Write()
}

package wsserver

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/handler"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	global.GVA_LOG.Info("WS server is running on", zap.String("address", addr))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		global.GVA_LOG.Error("WS server listenAndServe error", zap.String("error", err.Error()))
	}
}

func Start() {
	go run()
}

func serveWs(hub *socket.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		global.GVA_LOG.Error("WS server upgrade error", zap.String("error", err.Error()))
		return
	}
	client := socket.NewClient(hub, conn.RemoteAddr().String(), conn)
	go client.Read()
	go client.Write()
	hub.Connect <- client
}

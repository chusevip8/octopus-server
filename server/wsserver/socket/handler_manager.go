package socket

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"sync"
)

var (
	handlers     = make(map[uint]Handler)
	handlersLock sync.RWMutex
)

type Handler func(client *Client, data []byte)

func RegisterHandler(key uint, value Handler) {
	handlersLock.Lock()
	defer handlersLock.Unlock()
	handlers[key] = value
}

func GetHandler(key uint) (handler Handler, ok bool) {
	handlersLock.RLock()
	defer handlersLock.RUnlock()
	handler, ok = handlers[key]
	return
}

func RegisterAllHandlers() {
	RegisterHandler(protocol.CodeLogin, LoginHandler)
	RegisterHandler(protocol.CodeTaskRequest, TaskRequestHandler)
	RegisterHandler(protocol.CodeTaskStart, TaskStartHandler)
	RegisterHandler(protocol.CodeTaskFinish, TaskFinishHandler)
}

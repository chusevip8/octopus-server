package socket

import (
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
	RegisterHandler(Heartbeat, HeartbeatHandler)
	RegisterHandler(Login, LoginHandler)
}

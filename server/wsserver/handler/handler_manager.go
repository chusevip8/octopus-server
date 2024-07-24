package handler

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/protocol"
	"github.com/flipped-aurora/gin-vue-admin/server/wsserver/socket"
	"sync"
)

type Handler func(client *socket.Client, data []byte)

// HandlerManager 结构体
type HandlerManager struct {
	handlers     map[uint]Handler
	handlersLock sync.RWMutex
}

// NewHandlerManager 创建一个新的 HandlerManager 实例
func NewHandlerManager() *HandlerManager {
	return &HandlerManager{
		handlers: make(map[uint]Handler),
	}
}

// RegisterHandler 注册处理程序
func (hm *HandlerManager) registerHandler(key uint, value Handler) {
	hm.handlersLock.Lock()
	defer hm.handlersLock.Unlock()
	hm.handlers[key] = value
}

// GetHandler 获取处理程序
func (hm *HandlerManager) getHandler(key uint) (handler Handler, ok bool) {
	hm.handlersLock.RLock()
	defer hm.handlersLock.RUnlock()
	handler, ok = hm.handlers[key]
	return
}

// RegisterAllHandlers 注册所有处理程序
func (hm *HandlerManager) RegisterAllHandlers() {
	hm.registerHandler(protocol.CodeLogin, LoginHandler)
	hm.registerHandler(protocol.CodeTaskRequest, TaskRequestHandler)
	hm.registerHandler(protocol.CodeTaskStart, TaskStartHandler)
	hm.registerHandler(protocol.CodeTaskFinish, TaskFinishHandler)
}

func (hm *HandlerManager) HandleMessage(client *socket.Client, data []byte) {
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
	if handle, ok := hm.getHandler(message.Code); ok {
		handle(client, msgContent)
	} else {
		fmt.Println("Handler not found", client.Addr, "Code", message.Code)
	}
}

package socket

import "sync"

type ClientManager struct {
	Clients     map[uint]*Client
	ClientsLock sync.RWMutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		Clients: make(map[uint]*Client),
	}
}

func (clientManger *ClientManager) RemoveClient(client *Client) {
	clientManger.ClientsLock.Lock()
	defer clientManger.ClientsLock.Unlock()
	delete(clientManger.Clients, client.Id)
}

func (clientManger *ClientManager) AddClient(client *Client) {
	clientManger.ClientsLock.Lock()
	defer clientManger.ClientsLock.Unlock()
	clientManger.Clients[client.Id] = client
}

func (clientManger *ClientManager) GetClient(clientId uint) (client *Client, ok bool) {
	clientManger.ClientsLock.RLock()
	defer clientManger.ClientsLock.RUnlock()
	client, ok = clientManger.Clients[clientId]
	return
}

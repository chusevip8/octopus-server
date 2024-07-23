package socket

import "sync"

var (
	clients     = make(map[uint]*Client)
	clientsLock sync.RWMutex
)

func RemoveClient(client *Client) {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	delete(clients, client.Id)
}

func AddClient(client *Client) {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	clients[client.Id] = client
}

func GetClient(clientId uint) (client *Client, ok bool) {
	clientsLock.RLock()
	defer clientsLock.RUnlock()
	client, ok = clients[clientId]
	return
}

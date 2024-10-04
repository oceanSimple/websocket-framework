package wsFramework

import "sync"

type ClientMap struct {
	clients map[string]*Client
	mutex   sync.RWMutex
}

func (cm *ClientMap) AddClient(client *Client) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.clients[client.ID] = client
}

func (cm *ClientMap) RemoveClient(id string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	delete(cm.clients, id)
}

func (cm *ClientMap) GetClient(id string) *Client {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	return cm.clients[id]
}

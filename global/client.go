package global

import (
	"errors"
	wsFramework "github.com/oceanSimple/websocket-framework"
)

func GetClient(clientId string) (*wsFramework.Client, error) {
	ClientLock.RLock()
	defer ClientLock.RUnlock()
	if c, ok := Clients[clientId]; ok {
		return c, nil
	} else {
		return nil, errors.New("clientId not found")
	}
}

func AddClient(client *wsFramework.Client) {
	ClientLock.Lock()
	defer ClientLock.Unlock()
	Clients[client.ID] = client
}

func RemoveClient(clientId string) {
	ClientLock.Lock()
	defer ClientLock.Unlock()
	delete(Clients, clientId)
}

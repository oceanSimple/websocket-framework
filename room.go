package wsFramework

import "fmt"

type Room struct {
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"-"` // key: clientID, value: client
}

type RoomMap map[string]*Room

func (r RoomMap) getRoom(name string) (*Room, error) {
	if _, ok := r[name]; !ok {
		return nil, fmt.Errorf("Room %s not exists", name)
	}

	return r[name], nil
}

func (r RoomMap) newRoom(name string) error {
	if _, ok := r[name]; ok {
		return fmt.Errorf("Room %s already exists", name)
	}

	r[name] = &Room{
		Name:    name,
		Clients: make(map[string]*Client),
	}

	return nil
}

func (r RoomMap) addClientToRoom(roomName, clientID string, client *Client) error {
	if _, ok := r[roomName]; !ok {
		err := r.newRoom(roomName)
		if err != nil {
			return err
		}
	}

	r[roomName].Clients[clientID] = client

	return nil
}

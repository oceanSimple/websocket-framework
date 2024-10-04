package wsFramework

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID      string          `json:"id"`
	Context any             `json:"context"`
	Conn    *websocket.Conn `json:"-"` // websocket connection
	Rooms   []string        `json:"-"` // Room
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		ID:    id,
		Conn:  conn,
		Rooms: make([]string, 0),
	}
}

func (client *Client) Emit(ns string, event string, data interface{}) error {
	msg, err := marshalMessage(newMessage(ns, event, data))
	if err != nil {
		return err
	}
	return client.Conn.WriteMessage(websocket.TextMessage, msg)
}

func (client *Client) Join(roomName string) error {
	client.Rooms = append(client.Rooms, roomName)
	return rooms.addClientToRoom(roomName, client.ID, client)
}

func (client *Client) BroadcastToOthers(roomName string, nsName string, event string, data interface{}) error {
	return client.broadcast(false, roomName, nsName, event, data)
}

func (client *Client) BroadcastToAll(roomName string, nsName string, event string, data interface{}) error {
	return client.broadcast(true, roomName, nsName, event, data)
}

// flag: true broadcast to all clients
// flag: false broadcast to all clients except the client itself
func (client *Client) broadcast(flag bool, roomName string, nsName string, event string, data interface{}) error {
	// check room is existing
	room, err := rooms.getRoom(roomName)
	if err != nil {
		return err
	}
	// check client is in room
	if _, ok := room.Clients[client.ID]; !ok {
		return fmt.Errorf("client %s not in room %s", client.ID, roomName)
	}
	// broadcast to others
	for id, c := range room.Clients {
		if !flag {
			if id == client.ID {
				continue
			}
		}
		err := c.Emit(nsName, event, data)
		if err != nil {
			return err
		}
	}
	return nil
}

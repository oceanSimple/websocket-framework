package wsFramework

import (
	"github.com/gorilla/websocket"
	"github.com/oceanSimple/websocket-framework/class"
	"github.com/oceanSimple/websocket-framework/global"
)

type Client struct {
	ID      string          `json:"id"`
	Context any             `json:"context"`
	Conn    *websocket.Conn `json:"-"` // websocket connection
	Rooms   []string        `json:"-"` // room
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		ID:    id,
		Conn:  conn,
		Rooms: make([]string, 0),
	}
}

func (client *Client) Emit(ns string, event string, data interface{}) error {
	msg, err := class.MarshalMessage(class.NewMessage(ns, event, data))
	if err != nil {
		return err
	}
	return client.Conn.WriteMessage(websocket.TextMessage, msg)
}

func (client *Client) Join(roomName string) error {
	client.Rooms = append(client.Rooms, roomName)
	room, err := global.GetRoom(roomName)
	if err != nil {
		return err
	}

}

package model

import "github.com/gorilla/websocket"

type Client struct {
	ID      string          `json:"id"`
	Context any             `json:"context"`
	Conn    *websocket.Conn `json:"-"` // websocket connection
	Room    *Room           `json:"-"` // room
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		ID:   id,
		Conn: conn,
	}
}

package handler

import "github.com/gorilla/websocket"

type Handler interface {
	Handle([]byte, *websocket.Conn) error
}

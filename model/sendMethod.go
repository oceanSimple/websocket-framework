package model

import "github.com/gorilla/websocket"

type SendMethod func(conn *websocket.Conn, data any) error

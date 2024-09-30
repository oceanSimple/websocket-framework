package wsFramework

import (
	"fmt"
	"github.com/gorilla/websocket"
	"testing"
)

func TestRun(t *testing.T) {
	server := NewServer()
	server.On("test", "/test", func(conn *websocket.Conn, data any) error {
		fmt.Println("receive data: ", data)
		return conn.WriteMessage(websocket.TextMessage, []byte("hello"))
	})
	server.Run("9000")
}

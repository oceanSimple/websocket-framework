package wsFramework

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type ServerOption func(*Server)

type Server struct {
	Port string

	// upgrader is used to upgrade an HTTP connection to a WebSocket connection.
	upgrader *websocket.Upgrader
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{
		Port: "9000",

		upgrader: defaultUpgrader(),
	}
	for _, option := range options {
		option(server)
	}
	return server
}

func WithUpgrader(upgrader *websocket.Upgrader) ServerOption {
	return func(s *Server) {
		s.upgrader = upgrader
	}
}

func defaultUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

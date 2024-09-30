package wsFramework

import (
	"github.com/gorilla/websocket"
	"github.com/oceanSimple/websocket-framework/handler"
)

type ServerOption func(*Server)

type Server struct {
	Port string

	// upgrader is used to upgrade an HTTP connection to a WebSocket connection.
	upgrader *websocket.Upgrader

	handler handler.Handler
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{
		Port:     "9000",
		upgrader: defaultUpgrader(),
		handler:  handler.WsHandler{},
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

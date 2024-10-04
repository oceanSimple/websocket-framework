package wsFramework

import (
	"github.com/gorilla/websocket"
	"github.com/oceanSimple/websocket-framework/handler"
	"net/http"
)

type ServerOption func(*Server)

type Server struct {
	Port string

	// used to store all the namespace
	nameSpaces NameSpaceMap

	// used to store all the rooms
	rooms RoomMap

	// upgrader is used to upgrade an HTTP connection to a WebSocket connection.
	upgrader *websocket.Upgrader

	handler handler.Handler
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{
		Port: "9000",

		nameSpaces: make(map[string]*NameSpace),
		rooms:      make(map[string]*Room),

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

func defaultUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

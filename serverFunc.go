package wsFramework

import (
	"github.com/oceanSimple/websocket-framework/class"
	"github.com/oceanSimple/websocket-framework/global"
)

func (server *Server) On(ns string, path string, handler class.SendMethod) {
	// get the namespace
	namespace, err := global.GetNameSpace(ns)
	if err != nil {
		namespace = global.NewNameSpace(ns)
	}

	// add the handler
	namespace.Handlers[path] = handler
}

func (server *Server) Handle(bytes []byte, client *Client) error {
	message, err := parseMessage(bytes)
	if err != nil {
		return err
	}

	// get the namespace
	ns, err := server.nameSpaces.GetNameSpace(message.NameSpaceName)
	if err != nil {
		return err
	}

	// according to the path, get the handler
	handler, err := ns.GetHandler(message.Path)
	if err != nil {
		return err
	}

	return handler(client, message.Data)
}

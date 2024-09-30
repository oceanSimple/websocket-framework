package wsFramework

import (
	"github.com/oceanSimple/websocket-framework/global"
	"github.com/oceanSimple/websocket-framework/model"
)

func (server *Server) On(ns string, path string, handler model.SendMethod) {
	// get the namespace
	namespace, err := global.GetNameSpace(ns)
	if err != nil {
		namespace = global.NewNameSpace(ns)
	}

	// add the handler
	namespace.Handlers[path] = handler
}

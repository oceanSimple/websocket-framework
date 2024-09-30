package handler

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/oceanSimple/websocket-framework/global"
	"github.com/oceanSimple/websocket-framework/model"
)

type WsHandler struct{}

func (w WsHandler) Handle(bytes []byte, conn *websocket.Conn) error {
	// parse []byte to Message
	message, err := model.ParseMessage(bytes)
	if err != nil {
		return err
	}

	// get the namespace
	ns, err := global.GetNameSpace(message.NameSpaceName)
	if err != nil {
		return err
	}

	// according to the path, find the corresponding method
	method, ok := ns.Handlers[message.Path]
	if !ok {
		return errors.New("path not found")
	}

	// execute the method
	return method(conn, message.Data)
}

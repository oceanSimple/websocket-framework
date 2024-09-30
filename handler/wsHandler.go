package handler

import (
	"errors"
	"github.com/oceanSimple/websocket-framework"
	"github.com/oceanSimple/websocket-framework/class"
	"github.com/oceanSimple/websocket-framework/global"
)

type WsHandler struct{}

func (w WsHandler) Handle(bytes []byte, client *wsFramework.Client) error {
	// parse []byte to Message
	message, err := class.ParseMessage(bytes)
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
	return method(client, message.Data)
}

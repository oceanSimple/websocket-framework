package handler

import (
	"github.com/oceanSimple/websocket-framework"
)

type Handler interface {
	Handle([]byte, *wsFramework.Client) error
}

package class

import "github.com/oceanSimple/websocket-framework"

type Room struct {
	Name    string                         `json:"name"`
	Clients map[string]*wsFramework.Client `json:"-"` // key: clientID, value: client
}

package global

import (
	"github.com/oceanSimple/websocket-framework"
	"github.com/oceanSimple/websocket-framework/class"
	"sync"
)

var (
	NameSpaces = make(map[string]*class.NameSpace)
	Rooms      = make(map[string]*class.Room)
	Clients    = make(map[string]*wsFramework.Client)
)

var (
	NsRwLock   = new(sync.RWMutex)
	RoomRwLock = new(sync.RWMutex)
	ClientLock = new(sync.RWMutex)
)

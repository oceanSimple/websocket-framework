package global

import (
	"github.com/oceanSimple/websocket-framework/model"
	"sync"
)

var (
	NameSpaces = make(map[string]*model.NameSpace)
	Rooms      = make(map[string]*model.Room)
)

var (
	NsRwLock   = new(sync.RWMutex)
	RoomRwLock = new(sync.RWMutex)
)

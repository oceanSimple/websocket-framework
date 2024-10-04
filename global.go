package wsFramework

import "sync"

var (
	nameSpaces = make(NameSpaceMap)
	rooms      = make(RoomMap)
	clients    = &ClientMap{clients: make(map[string]*Client), mutex: sync.RWMutex{}}
)

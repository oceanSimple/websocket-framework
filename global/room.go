package global

import (
	"errors"
	"github.com/oceanSimple/websocket-framework/class"
)

func GetRoom(room string) (*class.Room, error) {
	RoomRwLock.RLock()
	defer RoomRwLock.RUnlock()
	if r, ok := Rooms[room]; ok {
		return r, nil
	} else {
		return nil, errors.New("room not found")
	}
}

func AddNewRoom(room string) error {
	RoomRwLock.Lock()
	defer RoomRwLock.Unlock()
	if _, ok := Rooms[room]; ok {
		return errors.New("room already exists")
	}
	Rooms[room] = &class.Room{
		Name:    room,
		Clients: make(map[string]*class.Client),
	}
	return nil
}

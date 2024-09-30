package global

import (
	"errors"
	"github.com/oceanSimple/websocket-framework/model"
)

func NewNameSpace(ns string) *model.NameSpace {
	nsModel := &model.NameSpace{
		Name:     ns,
		Handlers: make(map[string]model.SendMethod),
	}
	NameSpaces[ns] = nsModel
	return nsModel
}

func GetNameSpace(ns string) (*model.NameSpace, error) {
	NsRwLock.RLock()
	defer NsRwLock.RUnlock()
	if ns, ok := NameSpaces[ns]; ok {
		return ns, nil
	} else {
		return nil, errors.New("namespace not found")
	}
}

func GetRoom(room string) (*model.Room, error) {
	RoomRwLock.RLock()
	defer RoomRwLock.RUnlock()
	if r, ok := Rooms[room]; ok {
		return r, nil
	} else {
		return nil, errors.New("room not found")
	}
}

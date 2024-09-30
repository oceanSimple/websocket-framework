package global

import (
	"errors"
	"github.com/oceanSimple/websocket-framework/class"
)

func NewNameSpace(ns string) *class.NameSpace {
	nsModel := &class.NameSpace{
		Name:     ns,
		Handlers: make(map[string]class.SendMethod),
	}
	NameSpaces[ns] = nsModel
	return nsModel
}

func GetNameSpace(ns string) (*class.NameSpace, error) {
	NsRwLock.RLock()
	defer NsRwLock.RUnlock()
	if ns, ok := NameSpaces[ns]; ok {
		return ns, nil
	} else {
		return nil, errors.New("namespace not found")
	}
}

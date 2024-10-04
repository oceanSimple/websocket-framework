package wsFramework

import "fmt"

type NameSpace struct {
	Name     string                `json:"name"`
	Handlers map[string]SendMethod `json:"-"` // key: path, value: sendMethod
}

func (ns *NameSpace) GetHandler(path string) (SendMethod, error) {
	if _, ok := ns.Handlers[path]; !ok {
		return nil, fmt.Errorf("handler %s not exists", path)
	}
	return ns.Handlers[path], nil

}

type NameSpaceMap map[string]*NameSpace

func (m NameSpaceMap) GetNameSpace(name string) (*NameSpace, error) {
	if _, ok := m[name]; !ok {
		return nil, fmt.Errorf("NameSpace %s not exists", name)
	}

	return m[name], nil
}

func (m NameSpaceMap) newNameSpace(name string) (*NameSpace, error) {
	if _, ok := m[name]; ok {
		return nil, fmt.Errorf("NameSpace %s already exists", name)
	}

	m[name] = &NameSpace{
		Name:     name,
		Handlers: make(map[string]SendMethod),
	}

	return m[name], nil
}

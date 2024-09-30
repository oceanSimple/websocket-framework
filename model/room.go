package model

type Room struct {
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"-"` // key: clientID, value: client
}

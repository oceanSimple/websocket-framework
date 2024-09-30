package model

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	NameSpaceName string `json:"nameSpaceName"`
	Path          string `json:"path"`
	Data          any    `json:"data"`
}

func NewMessage(nameSpaceName, path string, data any) Message {
	return Message{
		NameSpaceName: nameSpaceName,
		Path:          path,
		Data:          data,
	}
}

func ParseMessage(message []byte) (Message, error) {
	var m Message
	err := json.Unmarshal(message, &m)
	if err != nil {
		return Message{}, fmt.Errorf("parse message error: %v", err)
	}
	return m, nil
}

func MarshalMessage(message Message) ([]byte, error) {
	m, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("marshal message error: %v", err)
	}
	return m, nil
}

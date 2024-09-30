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

func ParseMessage(message []byte) (Message, error) {
	var m Message
	err := json.Unmarshal(message, &m)
	if err != nil {
		return Message{}, fmt.Errorf("parse message error: %v", err)
	}
	return m, nil
}

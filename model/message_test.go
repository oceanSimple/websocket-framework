package model

import (
	"encoding/json"
	"testing"
)

func TestMessage(t *testing.T) {
	message := Message{
		NameSpaceName: "test",
		Path:          "/test",
		Data: Message{
			NameSpaceName: "ns",
			Path:          "path",
			Data:          "data",
		},
	}
	marshal, _ := json.Marshal(message)
	t.Log(string(marshal))
}

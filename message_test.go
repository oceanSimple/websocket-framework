package wsFramework

import (
	"fmt"
	"testing"
)

func TestJoinRoom(t *testing.T) {
	message := newMessage("test", "/joinRoom", "joinRoom data")
	bytes, _ := marshalMessage(message)
	fmt.Println(string(bytes))
}

func TestBroadcast(t *testing.T) {
	message := newMessage("test", "/broadcast", "broadcast data")
	bytes, _ := marshalMessage(message)
	fmt.Println(string(bytes))
}

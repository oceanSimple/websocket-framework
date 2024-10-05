package wsFramework

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	server := NewServer()
	server.On("test", "/test", func(client *Client, data any) error {
		fmt.Println("receive data: ", data)
		return client.Emit("test", "/test", "test data")
	})
	server.On("test", "/joinRoom", func(client *Client, data any) error {
		return client.Join("testRoom")
	})
	server.On("test", "/broadcast", func(client *Client, data any) error {
		return client.BroadcastToOthers("testRoom", "test", "/broadcast", "broadcast data")
	})
	server.Run("9000")
}

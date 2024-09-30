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
	server.Run("9000")
}

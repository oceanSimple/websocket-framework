package wsFramework

type SendMethod func(*Client, any) error

type Handler interface {
	Handle([]byte, *Client) error
}

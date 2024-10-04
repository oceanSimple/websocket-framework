package wsFramework

import "fmt"

func (server *Server) On(ns string, path string, handler SendMethod) {
	// get the namespace
	namespace, err := nameSpaces.GetNameSpace(ns)
	if err != nil {
		namespace, err = nameSpaces.newNameSpace(ns)
		if err != nil {
			fmt.Println(outputError() + err.Error())
			return
		}
	}

	// add the handler
	namespace.Handlers[path] = handler
}

func (server *Server) Handle(bytes []byte, client *Client) error {
	message, err := parseMessage(bytes)
	if err != nil {
		return err
	}

	// get the namespace
	ns, err := nameSpaces.GetNameSpace(message.NameSpaceName)
	if err != nil {
		return err
	}

	// according to the path, get the handler
	handler, err := ns.GetHandler(message.Path)
	if err != nil {
		return err
	}

	return handler(client, message.Data)
}

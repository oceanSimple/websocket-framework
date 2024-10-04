package wsFramework

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func (server *Server) Run(port string) {
	server.Port = port
	http.HandleFunc("/ws", server.handleWebSocket)
	log.Println(outputInfo() + "Websocket server started on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(outputError() + err.Error())
}

func (server *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrade http connection to websocket connection
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(outputError() + err.Error())
		return
	}

	// generate an uuid
	uid := uuid.New()
	uidMsg, _ := marshalMessage(newMessage("system", "/uuid", uid.String()))
	_ = conn.WriteMessage(websocket.TextMessage, uidMsg)

	defer func(conn *websocket.Conn) {
		// remove client from global clients
		clients.RemoveClient(uid.String())
		err := conn.Close()
		if err != nil {
			log.Println(outputError() + err.Error())
		}
	}(conn)

	// create a new client
	client := NewClient(uid.String(), conn)

	// add client to global clients
	clients.AddClient(client)

	// handle messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(outputError() + err.Error())
			return
		}

		// handle message
		err = server.Handle(msg, client)
		if err != nil {
			log.Println(outputError() + err.Error())
		}
	}
}

package wsFramework

import (
	"github.com/gorilla/websocket"
	"github.com/oceanSimple/websocket-framework/output"
	"log"
	"net/http"
)

func (server *Server) Run(port string) {
	server.Port = port
	http.HandleFunc("/ws", server.handleWebSocket)
	log.Println(output.OutputInfo() + "Server started on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(output.OutputError() + err.Error())
}

func (server *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(output.OutputError() + err.Error())
		}
	}(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(output.OutputError() + err.Error())
			return
		}
		log.Println("Received message:", string(msg))

		err = server.handler.Handle(msg, conn)
		if err != nil {
			log.Println(output.OutputError() + err.Error())
		}
	}
}

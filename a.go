package wsFramework

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/oceanSimple/websocket-framework/model"
	"github.com/oceanSimple/websocket-framework/output"
	"log"
	"net/http"
)

func (server *Server) Run(port string) {
	server.Port = port
	http.HandleFunc("/ws", server.handleWebSocket)
	log.Println(output.OutputInfo() + "Websocket server started on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(output.OutputError() + err.Error())
}

func (server *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrade http connection to websocket connection
	conn, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(output.OutputError() + err.Error())
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(output.OutputError() + err.Error())
		}
	}(conn)

	// generate an uuid
	uid := uuid.New()
	uidMsg, _ := model.MarshalMessage(model.NewMessage("system", "uuid", uid.String()))
	_ = conn.WriteMessage(websocket.TextMessage, uidMsg)

	// handle messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(output.OutputError() + err.Error())
			return
		}

		// handle message
		err = server.handler.Handle(msg, conn)
		if err != nil {
			log.Println(output.OutputError() + err.Error())
		}
	}
}

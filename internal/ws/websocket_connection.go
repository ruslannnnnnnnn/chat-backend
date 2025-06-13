package websocket_connection

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Received: %s ", message)

		if err := ws.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			break
		}
	}
}

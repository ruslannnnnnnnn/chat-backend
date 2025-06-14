package wsc

import (
	"log"

	"github.com/gorilla/websocket"
)

func HandleMessage(ws *websocket.Conn, message string, messageType int) {
	log.Printf("Received: %s ", message)

	modifiedMessage := "got: " + string(message)
	messageInBytes := []byte(modifiedMessage)

	if err := ws.WriteMessage(messageType, messageInBytes); err != nil {
		log.Println(err)
	}
}

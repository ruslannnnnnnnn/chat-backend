package wsc

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleChatConnections(w http.ResponseWriter, r *http.Request, hub *Hub) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	log.Println("подключился клиент", ws.RemoteAddr().String())

	go hub.Run()

	// Получаем имя пользователя из параметров запроса
	queryParams := r.URL.Query()
	name := "unknown"
	if queryParams.Has("name") {
		name = queryParams.Get("name")
	}

	client := &Client{
		Name: name,
		conn: ws,
	}

	// добавляем клиента
	hub.ClientPipe <- client

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("Пришло сообщение", string(message))
		//отправляем сообщения
		hub.MessagePipe <- &Message{
			MessageType: messageType,
			Message:     message,
			Sender:      client,
		}

	}
}

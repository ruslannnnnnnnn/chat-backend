package wsc

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleChatConnections(w http.ResponseWriter, r *http.Request, hub *Hub) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
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

	clientUuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("Не удалось сгенерировать uuid")
	}

	client := &Client{
		Id:   clientUuid.String(),
		Name: name,
		conn: ws,
	}

	// добавляем клиента
	hub.ClientPipe <- client

	defer func() {
		hub.DisconnectClientPipe <- client.Id
		log.Println("Клиент отключился id:", client.Id, "ip:", client.conn.RemoteAddr())
	}()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println("Соединение закрыто с ошибкой, id клиента:", client.Id, "ошибка:", err.Error())
			} else {
				log.Println("Соединение закрыто, id клиента: ", client.Id)
			}
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

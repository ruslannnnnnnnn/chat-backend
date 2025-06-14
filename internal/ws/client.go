package wsc

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Name string
	conn *websocket.Conn
}

func (client *Client) Send(message *Message) {

	jsonMessageBody := JsonMessageBody{
		message.Sender.Name,
		string(message.Message),
	}
	serialized, err := json.Marshal(jsonMessageBody)
	if err != nil {
		log.Println("Не удалось создать жсонку")
	}

	err = client.conn.WriteMessage(message.MessageType, serialized)
	if err != nil {
		log.Fatal("Не удалось отправить сообщение клиенту")
	}

	log.Println("Отправили сообщение клиенту", client.conn.RemoteAddr().String())
}

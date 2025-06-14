package wsc

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func (client *Client) Send(message *Message) {
	log.Println("Отправили сообщение клиенту", client.conn.RemoteAddr().String())
	err := client.conn.WriteMessage(message.MessageType, message.Message)
	if err != nil {
		log.Fatal("Не удалось отправить сообщение клиенту")
	}
}

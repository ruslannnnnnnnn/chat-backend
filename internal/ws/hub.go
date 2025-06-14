package wsc

import "log"

type Hub struct {
	Clients     []*Client
	MessagePipe chan *Message
	ClientPipe  chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:     make([]*Client, 0),
		MessagePipe: make(chan *Message),
		ClientPipe:  make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		//когда приходит новое сообщение, рассылаем всем клиентам
		case message := <-h.MessagePipe:
			log.Println("Принимаем сообщение в селекте", string(message.Message))
			for _, client := range h.Clients {
				go client.Send(message)
			}
		//когда подключается новый клиент, добавляем в слайс с клиентами
		case newClient := <-h.ClientPipe:
			log.Println("Добавили клиента в хаб", newClient.conn.RemoteAddr().String())
			h.Clients = append(h.Clients, newClient)
		}

	}
}

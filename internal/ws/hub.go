package wsc

import (
	"log"
	"sync"
)

type Hub struct {
	mu                   sync.Mutex
	Clients              map[string]*Client
	MessagePipe          chan *Message
	ClientPipe           chan *Client
	DisconnectClientPipe chan string
}

func NewHub() *Hub {
	return &Hub{
		Clients:              make(map[string]*Client),
		MessagePipe:          make(chan *Message),
		ClientPipe:           make(chan *Client),
		DisconnectClientPipe: make(chan string),
	}
}

func (h *Hub) Run() {
	for {
		select {
		//когда приходит новое сообщение, рассылаем всем клиентам
		case message := <-h.MessagePipe:
			log.Println("Принимаем сообщение в селекте", string(message.Message))

			h.mu.Lock()
			lockedClients := make([]*Client, 0, len(h.Clients))
			for _, client := range h.Clients {
				lockedClients = append(lockedClients, client)
			}
			h.mu.Unlock()

			for _, lockedClient := range lockedClients {
				go lockedClient.Send(message)
			}

		// когда подключается новый клиент, добавляем в мапу с клиентами
		case newClient := <-h.ClientPipe:
			h.mu.Lock()
			h.Clients[newClient.Id] = newClient
			h.mu.Unlock()
			log.Println("Добавили клиента в хаб, ip:", newClient.conn.RemoteAddr().String(), "uuid:", newClient.Id)
		// при отключении клиента удаляем его из мапы с клиентами
		case clientId := <-h.DisconnectClientPipe:
			h.mu.Lock()
			delete(h.Clients, clientId)
			h.mu.Unlock()
			log.Println("Удалили клиента из хаба", clientId)
		}
	}
}

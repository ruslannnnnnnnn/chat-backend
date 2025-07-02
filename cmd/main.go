package main

import (
	wsc "chat/internal/ws"
	"log"
	"net/http"
)

func main() {

	hub := wsc.NewHub()

	http.HandleFunc("/api/ws/chat", func(w http.ResponseWriter, r *http.Request) {
		wsc.HandleChatConnections(w, r, hub)
	})

	log.Println(("http server started on :8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

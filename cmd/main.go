package main

import (
	wsc "chat/internal/ws"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/ws", wsc.HandleConnections)
	log.Println(("http server started on :8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

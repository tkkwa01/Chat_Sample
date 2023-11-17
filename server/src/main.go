package main

import (
	"fmt"
	"github.com/tkkwa01/Chat_Sample/handlers"
	"github.com/tkkwa01/Chat_Sample/src/domain"
	"log"
	"net/http"
)

func main() {
	hub := domain.NewHub()
	go hub.RunLoop()

	http.HandleFunc("/ws", handlers.NewWebsocketHandler(hub).Handle)

	port := "80"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}

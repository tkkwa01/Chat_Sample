package main

import (
	"fmt"
	"github.com/tkkwa01/Chat_Sample/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handlers.NewWebsocketHandler().Handle)

	port := "80"
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Panicln("Serve Error:", err)
	}
}

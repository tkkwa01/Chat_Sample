package main

import (
	"fmt"
	"go/Chat_Sample/server/handlers"
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

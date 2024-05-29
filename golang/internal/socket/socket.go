package socket

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/", handleConnections)
	go handleMessages()
	// log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

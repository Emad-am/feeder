package socketservice

import (
	"log"
	"net"
)

var conn net.Conn
var list []net.Conn

func StartService() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server is listening on localhost:8080")

	list = []net.Conn{}
	for {
		conn, err = ln.Accept()
		list = append(list, conn)
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		log.Println("Accepted new connection")

		go handleConnection(conn)
	}
}

package socketservice

import (
	"fmt"
	"io"
	"log"
	"net"
	"slices"
	"time"

	"github.com/Emad-am/feeder/internal/socket"
	quotemakerservice "github.com/Emad-am/feeder/src/services/quote_maker"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func handleConnection(conn net.Conn) {
	defer func() {
		index := slices.IndexFunc(list, func(c net.Conn) bool { return c == conn })
		list = slices.Delete(list, index, index+1)
		conn.Close()
	}()
	log.Println("New WebSocket connection established")

	// Upgrade the connection to WebSocket
	hs, err := ws.Upgrade(conn)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}

	log.Printf("Handshake: %+v\n", hs)

	go func() {
		ch := quotemakerservice.GetChannel()
		for {
			message := <-*ch
			for _, conn := range list {
				SendMessage(conn, message)
			}
		}
	}()

	for {
		header, err := ws.ReadHeader(conn)
		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected")
				return
			}
			log.Println("Error reading header:", err)
			return
		}

		payload := make([]byte, header.Length)
		_, err = io.ReadFull(conn, payload)
		if err != nil {
			log.Println("Error reading payload:", err)
			return
		}

		if header.Masked {
			ws.Cipher(payload, header.Mask, 0)
		}

		// Reset the Masked flag, server frames must not be masked as per RFC6455
		header.Masked = false

		if err := ws.WriteHeader(conn, header); err != nil {
			log.Println("Error writing header:", err)
			return
		}

		if _, err := conn.Write(payload); err != nil {
			log.Println("Error writing payload:", err)
			return
		}

		if header.OpCode == ws.OpClose {
			log.Println("WebSocket connection closed")
			return
		}
	}
}
func SendMessage(conn net.Conn, message socket.Message) {
	text := fmt.Sprintf("%s %d %d %d %d %d %d",
		message.Symbol,
		message.Value.Ask,
		message.Value.Bid,
		message.Value.High,
		message.Value.Low,
		message.Value.Volume,
		time.Now().Unix(),
	)
	err := wsutil.WriteServerText(conn, []byte(text))
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}
	log.Println("Sent message to client:", message)
}

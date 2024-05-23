package feederservice

import (
	"fmt"
	"os"
	"time"
)

func startKeepAlive() {
	go func() {
		for {
			time.Sleep(10 * time.Second)
			_, err := fmt.Fprintf(conn, "KEEP_ALIVE\n")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error sending keep-alive message: %v\n", err)
				return
			}
			fmt.Println("Sent keep-alive message")
		}
	}()
}

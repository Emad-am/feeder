package feederservice

import (
	"fmt"
	"os"
	"time"
)

func (feeder *Feeder) startKeepAlive() {
	// fmt.Println("Starting keep-alive...")
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for range ticker.C {
			_, err := fmt.Fprintf(*feeder.Conn, "KEEP_ALIVE\n")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error sending keep-alive message: %v\n", err)
				return
			}
			fmt.Println("Sent keep-alive message")
		}
	}()
}

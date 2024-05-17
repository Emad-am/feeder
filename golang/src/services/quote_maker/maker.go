package quotemakerservice

import (
	feederservice "dde/src/services/feeder"
	"fmt"
	"time"
)

func Start() {
	symbols := feederservice.GetSymbols()

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			for symbol, val := range *symbols {
				fmt.Println(symbol, val)
			}
		}
	}()
}

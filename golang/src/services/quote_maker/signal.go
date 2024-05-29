package quotemakerservice

import (
	"time"

	feederservice "github.com/Emad-am/feeder/src/services/feeder"
)

func startSignalSender(c *chan feederservice.Tick) {
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			go func() {
				*c <- feederservice.Tick{
					Symbol: "make-quote",
				}
			}()
		}
	}()
}

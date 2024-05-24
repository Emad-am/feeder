package quotemakerservice

import (
	"fmt"
	"time"

	"github.com/Emad-am/feeder/internal/redis"
	feederservice "github.com/Emad-am/feeder/src/services/feeder"
)

func Start() {
	c := *feederservice.GetChannel()
	symbs := *feederservice.NewSymbols()

	go func() {

		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			go func() {
				c <- feederservice.Tick{
					Symbol: "make-quote",
				}
			}()
		}
	}()

	go func() {

		for {

			tick := <-c

			if tick.Symbol == "make-quote" {

				for symbol, val := range symbs {
					if val.Volume != 0 {
						redis.Lpush(symbol, val)
						fmt.Println(symbol, val)
					}
				}
				symbs = *feederservice.NewSymbols()
				continue
			}

			symbol := symbs[tick.Symbol]
			// if symbol.Ask < tick.Ask {
			symbol.Ask = tick.Ask
			// }
			// if symbol.Bid < tick.Bid {
			symbol.Bid = tick.Bid
			//			}
			if symbol.High < tick.Bid {
				symbol.High = tick.Bid
			}
			symbol.Time = tick.Time

			if symbol.Low != 0 {
				if symbol.Low > tick.Bid {
					symbol.Low = tick.Bid
				}
			} else {
				symbol.Low = tick.Bid
			}

			symbol.Volume += 1

			// TODO : pointer
			symbs[tick.Symbol] = symbol
		}
	}()

}

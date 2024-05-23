package quotemakerservice

import (
	"fmt"
	"sync"
	"time"

	"github.com/Emad-am/feeder/internal/redis"
	feederservice "github.com/Emad-am/feeder/src/services/feeder"
)

var mu sync.Mutex

func Start() {
	c := *feederservice.GetChannel()
	symbs := *feederservice.NewSymbols()
	go func() {
		for {
			mu.Lock()
			tick := <-c
			symbol := symbs[tick.Symbol]
			if symbol.Ask < tick.Ask {
				symbol.Ask = tick.Ask
			}
			if symbol.Bid < tick.Bid {
				symbol.Bid = tick.Bid
			}
			if symbol.High < tick.Ask {
				symbol.High = tick.Ask
			}
			if symbol.Time < tick.Time {
				symbol.Time = tick.Time
			}
			if symbol.Low != 0 {
				if symbol.Low > tick.Ask {
					symbol.Low = tick.Ask
				}
			} else {
				symbol.Low = tick.Ask
			}

			symbol.Volume += 1

			symbs[tick.Symbol] = symbol
			mu.Unlock()
		}
	}()

	go func() {
		for {
			mu.Lock()
			for symbol, val := range symbs {
				if val.Volume != 0 {
					redis.Lpush(symbol, val)
					fmt.Println(symbol, val)
				}
			}
			mu.Unlock()
			symbs = *feederservice.NewSymbols()
			time.Sleep(500 * time.Millisecond)
		}
	}()
}

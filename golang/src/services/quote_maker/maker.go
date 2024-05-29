package quotemakerservice

import (
	"github.com/Emad-am/feeder/internal/redis"
	"github.com/Emad-am/feeder/internal/socket"
	feederservice "github.com/Emad-am/feeder/src/services/feeder"
	"github.com/Emad-am/feeder/tools"
)

var Chnn = make(chan socket.Message)

func GetChannel() *chan socket.Message { return &Chnn }

func startMaker(feeder *feederservice.Feeder) {
	c := *feeder.C
	symbs := *feeder.Symbols
	go func() {
		for {
			tick := <-c
			if tick.Symbol == "make-quote" {
				for symbol, val := range symbs {
					if val.Volume != 0 {
						x := socket.Message{
							Symbol: symbol,
							Value:  val,
						}
						go func(data socket.Message) {
							Chnn <- data
						}(x)

						// fmt.Println(symbol, val)
						go redis.Hset("symbols", symbol, tools.StructToString(val))
					}
					go redis.Publish("symbols-channel", tools.StructToString(symbs))
				}
				// symbs = *feederservice.NewSymbols()
				symbs = *feederservice.NewSymbols()
				continue
			}
			symbol := symbs[tick.Symbol]

			symbol.Ask = tick.Ask
			symbol.Bid = tick.Bid
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

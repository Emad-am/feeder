package main

import (
	_ "net/http/pprof"

	_ "github.com/Emad-am/feeder/internal/context"
	feederservice "github.com/Emad-am/feeder/src/services/feeder"
	quotemakerservice "github.com/Emad-am/feeder/src/services/quote_maker"
	socketservice "github.com/Emad-am/feeder/src/services/socket"
)

func main() {
	// defer func() {
	// 	db.Disconnect()
	// }()

	// go func() {
	// 	fmt.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	feeder := feederservice.NewFeeder()
	quotemakerservice.Start(feeder)
	feeder.Start()
	socketservice.StartService()
}

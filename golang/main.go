package main

import (
	_ "dde/internal/context"
	feederservice "dde/src/services/feeder"
	quotemakerservice "dde/src/services/quote_maker"
)

func main() {
	// defer func() {
	// 	db.Disconnect()
	// }()

	quotemakerservice.Start()
	feederservice.Start()

}

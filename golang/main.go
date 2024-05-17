package main

import (
	_ "dde/internal/context"
	feederservice "dde/src/services/feeder"
)

func main() {
	// defer func() {
	// 	db.Disconnect()
	// }()

	feederservice.Start()

}

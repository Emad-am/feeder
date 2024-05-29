package quotemakerservice

import (
	feederservice "github.com/Emad-am/feeder/src/services/feeder"
)

func Start(feeder *feederservice.Feeder) {
	startSignalSender(feeder.C)
	startMaker(feeder)
}

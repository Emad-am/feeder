package feederservice

import (
	"fmt"
	"net"
	"os"

	"github.com/Emad-am/feeder/src/services/api"
)

type Feeder struct {
	api     *api.Api
	C       *(chan Tick)
	Symbols *symbols
	Conn    *net.Conn
}

func (f *Feeder) Start() {
	f.startKeepAlive()
	f.startReader()

	f.api.Start()
}

func NewFeeder(api *api.Api) *Feeder {
	c := make(chan Tick, 100)
	host := conf.DDE.Host
	port := conf.DDE.Port
	connString := fmt.Sprintf("%s:%s", host, port)
	conn, e := net.Dial("tcp", connString)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to server: %v\n", e)
		return nil
	}
	return &Feeder{
		C:       &c,
		Symbols: NewSymbols(),
		Conn:    &conn,
	}
}

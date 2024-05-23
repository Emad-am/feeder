package feederservice

import (
	"sync"

	"github.com/Emad-am/feeder/config"
)

var conf = config.GetConfig()
var syms symbols = *NewSymbols()
var c = make(chan tick, 1)
var wg sync.WaitGroup

type tick struct {
	Symbol string
	Ask    int64
	Bid    int64
	Time   int64
}

type quote struct {
	Ask    int64
	Bid    int64
	High   int64
	Low    int64
	Volume int64
	Time   int64
}

type symbols map[string]quote

func NewSymbols() *symbols {
	return &symbols{
		"AUDCAD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDCHF": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDJPY": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDNZD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"BCHUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"BTCEUR": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"BTCUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"CADCHF": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"CADJPY": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"CHFJPY": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"ETHUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURAUD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURCAD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURCHF": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURGBP": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURJPY": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURNZD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURPLN": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURSEK": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPAUD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPCAD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPCHF": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPJPY": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPNZD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GOLD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GOLDEURO": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"LTCUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDCAD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDCHF": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDJPY": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"Palladium": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"Platinum": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"SILVER": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"SILVEREURO": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDCAD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDCHF": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDDKK": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDHKD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDJPY": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDMXN": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDNOK": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDPLN": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDSEK": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDSGD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDZAR": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"XRPUSD": quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
	}
}

func GetSymbols() *symbols {
	return &syms
}

func GetChannel() *chan tick {
	return &c
}

func GetWaitGroup() *sync.WaitGroup {
	return &wg
}

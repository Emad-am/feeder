package feederservice

import (
	"github.com/Emad-am/feeder/config"
)

var conf = config.GetConfig()

type Tick struct {
	Symbol string
	Ask    int64
	Bid    int64
	Time   int64
}

type Quote struct {
	Ask    int64
	Bid    int64
	High   int64
	Low    int64
	Volume int64
	Time   int64
}

type symbols map[string]Quote

func NewSymbols() *symbols {
	return &symbols{
		"AUDCAD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDCHF": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDJPY": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDNZD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"AUDUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"BCHUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"BTCEUR": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"BTCUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"CADCHF": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"CADJPY": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"CHFJPY": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"ETHUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURAUD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURCAD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURCHF": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURGBP": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURJPY": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURNZD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURPLN": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURSEK": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"EURUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPAUD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPCAD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPCHF": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPJPY": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPNZD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GBPUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GOLD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"GOLDEURO": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"LTCUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDCAD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDCHF": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDJPY": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"NZDUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"Palladium": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"Platinum": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"SILVER": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"SILVEREURO": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDCAD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDCHF": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDDKK": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDHKD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDJPY": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDMXN": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDNOK": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDPLN": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDSEK": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDSGD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"USDZAR": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
		"XRPUSD": Quote{
			Ask:    0,
			Bid:    0,
			High:   0,
			Low:    0,
			Volume: 0,
			Time:   0,
		},
	}
}

// func GetSymbols() *symbols {
// 	return &syms
// }

// func GetChannel() *chan Tick {
// 	return &c
// }

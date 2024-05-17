package feederservice

import "time"

type tick struct {
	Ask  string
	Bid  string
	Time time.Time
}
type ticks []tick
type symbols map[string]ticks

func NewSymbols() symbols {
	return symbols{
		"AUDCAD":     ticks{},
		"AUDCHF":     ticks{},
		"AUDJPY":     ticks{},
		"AUDNZD":     ticks{},
		"AUDUSD":     ticks{},
		"BCHUSD":     ticks{},
		"BTCEUR":     ticks{},
		"BTCUSD":     ticks{},
		"CADCHF":     ticks{},
		"CADJPY":     ticks{},
		"CHFJPY":     ticks{},
		"ETHUSD":     ticks{},
		"EURAUD":     ticks{},
		"EURCAD":     ticks{},
		"EURCHF":     ticks{},
		"EURGBP":     ticks{},
		"EURJPY":     ticks{},
		"EURNZD":     ticks{},
		"EURPLN":     ticks{},
		"EURSEK":     ticks{},
		"EURUSD":     ticks{},
		"GBPAUD":     ticks{},
		"GBPCAD":     ticks{},
		"GBPCHF":     ticks{},
		"GBPJPY":     ticks{},
		"GBPNZD":     ticks{},
		"GBPUSD":     ticks{},
		"GOLD":       ticks{},
		"GOLDEURO":   ticks{},
		"LTCUSD":     ticks{},
		"NZDCAD":     ticks{},
		"NZDCHF":     ticks{},
		"NZDJPY":     ticks{},
		"NZDUSD":     ticks{},
		"Palladium":  ticks{},
		"Platinum":   ticks{},
		"SILVER":     ticks{},
		"SILVEREURO": ticks{},
		"USDCAD":     ticks{},
		"USDCHF":     ticks{},
		"USDDKK":     ticks{},
		"USDHKD":     ticks{},
		"USDJPY":     ticks{},
		"USDMXN":     ticks{},
		"USDNOK":     ticks{},
		"USDPLN":     ticks{},
		"USDSEK":     ticks{},
		"USDSGD":     ticks{},
		"USDZAR":     ticks{},
		"XRPUSD":     ticks{},
	}
}

package barmodel

import (
	"time"

	"github.com/Emad-am/feeder/internal/db"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type BarModel struct {
	High   int64
	Low    int64
	Open   int64
	Close  int64
	Volume int64
	Time   time.Time
}

func NewBar(ticks []string) *BarModel {
	bar := BarModel{}
	return &bar
}

func NewBarPoint(symbol string, tf string, bar BarModel) *write.Point {
	point := influxdb2.NewPoint(
		"bar",
		map[string]string{
			"symbol":    symbol,
			"timeframe": tf,
		},
		map[string]interface{}{
			"high":   bar.High,
			"low":    bar.Low,
			"open":   bar.Open,
			"close":  bar.Close,
			"volume": bar.Volume,
		},
		bar.Time,
	)

	return point
}

func Save(p *write.Point) {
	db.GetWriteApi().WritePoint(p)
}

package db

import (
	"fmt"
	"log"

	"dde/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var client influxdb2.Client
var cnfg = config.GetConfig()

func init() {
	host := cnfg.Db.Host
	port := cnfg.Db.Port
	authToken := cnfg.Db.AuthToken

	connectionString := fmt.Sprintf("http://%s:%s", host, port)
	client = influxdb2.NewClient(connectionString, authToken)
	log.Println("Successfully connected to database")
}

func Disconnect() {
	client.Close()
}

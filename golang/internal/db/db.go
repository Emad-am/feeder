package db

import (
	"fmt"
	"log"

	"dde/config"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type influxDBClient struct {
	Client *influxdb2.Client
}

var client influxdb2.Client

func init() {

	cnfg := config.GetConfig()

	host := cnfg.Db.Host
	port := cnfg.Db.Port
	authToken := cnfg.Db.AuthToken

	connectionString := fmt.Sprintf("http://%s:%s", host, port)
	client = influxdb2.NewClient(connectionString, authToken)
	log.Println("Successfully connected to database")
}

func GetInfluxDBClient() *influxDBClient {
	return &influxDBClient{
		Client: &client,
	}
}

func Disconnect() {
	client.Close()
}

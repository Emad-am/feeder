package db

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func GetInfluxDBClient() *influxdb2.Client {
	return &client
}

func GetQueryApi() api.QueryAPI {
	return client.QueryAPI(cnfg.Db.Org)
}

func GetWriteApi() api.WriteAPI {
	return client.WriteAPI(cnfg.Db.Org, cnfg.Db.Bucket)
}



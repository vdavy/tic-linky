package influxdb

import (
	"github.com/sirupsen/logrus"
	"tic-linky/internal/config"
)
import "github.com/influxdata/influxdb/client/v2"

var logger = logrus.WithField("logger", "processing")

// influxDBClient the influxdb client
var influxDBClient client.Client

// NewInfluxdbClient create the connection to influxdb
func NewInfluxdbClient() {
	var err error
	if influxDBClient, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.InfluxdbServerURL,
		Username: config.InfluxdbUsername,
		Password: config.InfluxdbPassword,
	}); err != nil {
		logger.WithError(err).Fatal("Error connecting to Influxdb")
	}
}

// CloseInfluxDBConnection close the influxdb connection
func CloseInfluxDBConnection() {
	if err := influxDBClient.Close(); err != nil {
		logger.WithError(err).Error("Error closing connection to Influxdb")
	}
}

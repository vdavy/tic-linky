package influxdb

import (
	"github.com/influxdata/influxdb/client/v2"
	"tic-linky/internal/config"
)

// WritePointsIntoInflux write points into Influx
func WritePointsIntoInflux(influxdbPoints ...*client.Point) bool {
	batchPoint, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: config.InfluxdbDatabase,
	})
	if err != nil {
		logger.WithError(err).Error("Error creating batch point for Influx")
		return false
	}

	batchPoint.AddPoints(influxdbPoints)
	if err := influxDBClient.Write(batchPoint); err != nil {
		logger.WithError(err).Error("Error inserting points into InfluxDB")
		return false
	}

	return true
}

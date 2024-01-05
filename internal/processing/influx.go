package processing

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/sirupsen/logrus"
	"tic-linky/internal/influxdb"
)

// processEndOfFrame send the collected data to influxdb
func (frameData *frameData) processEndOfFrame() {
	if !frameData.checkDataValidity() {
		return
	}

	// generate influxdb points
	var pointList []*client.Point
	frameData.generateIndexesPoint(&pointList)
	frameData.generatePowerPoint(&pointList)
	frameData.generateDatedDataPoints(&pointList)
	frameData.generateSTGEPoint(&pointList)

	if logrus.IsLevelEnabled(logrus.DebugLevel) {
		logger.Debugf("%d points to insert into Influxdb :", len(pointList))
		spew.Dump(pointList)
	}

	// write the points
	pointsInserted := influxdb.WritePointsIntoInflux(pointList...)

	// flag dated points as inserted if done
	if pointsInserted {
		for fieldName, flaggedToSend := range frameData.datedFieldsWriteFlagMap {
			if flaggedToSend {
				frameData.datedFieldsWriteFlagMap[fieldName] = false
			}
		}
	}
}

// checkDataValidity check the extracted data from TIC are valid to insert into Influxdb
func (frameData *frameData) checkDataValidity() bool {
	if frameData.date == nil {
		logger.Warn("Data date is nil")
		return false
	}
	if len(frameData.indexesMap) == 0 {
		logger.Warn("No index data")
		return false
	}
	if len(frameData.powerMap) == 0 {
		logger.Warn("No power data")
		return false
	}
	if len(frameData.datedFieldsMap) == 0 {
		logger.Warn("No dated data")
		return false
	}

	return true
}

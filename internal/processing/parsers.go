package processing

import (
	"strconv"
	"time"
)

const dateValueFormat = "060102150405"

// parseDate parse the data field (used for influxdb point date)
func (frameData *frameData) parseDate(dataValue string) {
	frameDate, err := time.Parse(dateValueFormat, dataValue[1:])
	if err != nil {
		logger.WithError(err).Warnf("Error parsing frame date %s", dataValue)
		return
	}

	frameData.date = &frameDate
}

// parseFieldAsUint64 parse and store field and field value into map
func parseFieldAsUint64(fieldMap map[string]uint64, fieldName, fieldValue string) {
	fieldValueAsInt, err := strconv.ParseUint(fieldValue, 10, 64)
	if err != nil {
		logger.WithError(err).Warnf("Error converting field value %s for field name %s", fieldValue, fieldName)
		return
	}

	fieldMap[fieldName] = fieldValueAsInt
}

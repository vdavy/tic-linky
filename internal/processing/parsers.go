package processing

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"time"
)

const (
	dateValueFormat    = "060102150405"
	productionOffSet   = 10
	distributionOffSet = 14
)

// parseDate parse date field (used for influxdb point date)
func parseDate(dateField **time.Time, dataValue string) {
	frameDate, err := time.Parse(dateValueFormat, dataValue[1:])
	if err != nil {
		logger.WithError(err).Warnf("Error parsing frame date %s", dataValue)
		return
	}

	*dateField = &frameDate
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

// parseDatedField parse and store dated field value, flagging it to be sent
func (frameData *frameData) parseDatedField(line []string) {
	fieldName := line[fieldNameIndex]
	var fieldDate *time.Time
	parseDate(&fieldDate, line[dateFieldIndex])

	if fieldDate != nil { // date properly parsed
		if frameData.datedFieldsMap[fieldName].date == nil || // we don't have a date
			(frameData.datedFieldsMap[fieldName].date != nil && *frameData.datedFieldsMap[fieldName].date != *fieldDate) { // or the saved date is different from parsed date
			fieldValueAsInt, err := strconv.ParseUint(line[valueFieldIndex], 10, 64)
			if err != nil {
				logger.WithError(err).Warnf("Error converting field value %s for dated field name %s", line[valueFieldIndex], fieldName)
				return
			}

			// save the new data
			frameData.datedFieldsMap[fieldName] = datedField{
				date:  fieldDate,
				value: fieldValueAsInt,
			}
			frameData.datedFieldsWriteFlagMap[fieldName] = true // flag it to be sent
		}
	} else {
		logger.Warnf("Nil value for date in data field %s", fieldDate)
	}
}

func (frameData *frameData) parseSTGE(dataValue string) {
	hexBytes, err := hex.DecodeString(dataValue)
	if err != nil {
		logger.WithError(err).Warnf("Error decoding STGE field value %s", dataValue)
		return
	}
	stgeValue := binary.BigEndian.Uint32(hexBytes)
	frameData.productionIndex = int((stgeValue&(0xF<<productionOffSet))>>productionOffSet) + 1
	frameData.distributionIndex = int((stgeValue&(0x3<<distributionOffSet))>>distributionOffSet) + 1
}

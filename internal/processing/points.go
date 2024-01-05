package processing

import (
	"github.com/influxdata/influxdb/client/v2"
	"strconv"
	"tic-linky/internal/influxdb"
)

// generateIndexesPoint generate the Influxdb point from indexes data
func (frameData *frameData) generateIndexesPoint(pointList *[]*client.Point) {
	indexesPoint, err := client.NewPoint(influxdb.IndexesMeasurement,
		map[string]string{
			influxdb.ProductionTag:   intToString(frameData.productionIndex),
			influxdb.DistributionTag: intToString(frameData.distributionIndex),
		},
		convertTypedMapToAnyMap(frameData.indexesMap),
		*frameData.date)
	if err != nil {
		logger.WithError(err).Warn("Error generating indexes Influxdb point")
		return
	}

	*pointList = append(*pointList, indexesPoint)
}

// generatePowerPoint generate the Influxdb point from power data
func (frameData *frameData) generatePowerPoint(pointList *[]*client.Point) {
	powerPoint, err := client.NewPoint(influxdb.PowerMeasurement, nil,
		convertTypedMapToAnyMap(frameData.powerMap),
		*frameData.date)
	if err != nil {
		logger.WithError(err).Warn("Error generating power Influxdb point")
		return
	}

	*pointList = append(*pointList, powerPoint)
}

// generateDatedDataPoints generate a Influxdb point for each dated data
func (frameData *frameData) generateDatedDataPoints(pointList *[]*client.Point) {
	for fieldName, flaggedToSend := range frameData.datedFieldsWriteFlagMap { // process each dated data
		datedData := frameData.datedFieldsMap[fieldName] // get the extracted values

		if flaggedToSend && datedData.date != nil { // send only if needed and we have a date
			dateDataPoint, err := client.NewPoint(fieldName, nil,
				map[string]interface{}{
					influxdb.DatedDataFieldName: datedData.value,
				},
				*datedData.date)
			if err != nil {
				logger.WithError(err).Warnf("Error generating dated data Influxdb point for %s", fieldName)
				continue // skip to the next one
			}

			// add the generated to point to the list
			*pointList = append(*pointList, dateDataPoint)
		}
	}
}

// generateSTGEPoint generate the STGE point with data and the message field
func (frameData *frameData) generateSTGEPoint(pointList *[]*client.Point) {
	stgePoint, err := client.NewPoint(stgeField, nil,
		map[string]interface{}{
			influxdb.ContactSecOuvertField:             frameData.contactSecOuvertFlag,
			influxdb.OrganeDeCoupureField:              frameData.organeDeCoupureState,
			influxdb.CacheBorneDistributeurOuvertField: frameData.cacheBorneDistributeurOuvertFlag,
			influxdb.SurtensionField:                   frameData.surtensionFlag,
			influxdb.DepassementPuissanceField:         frameData.depassementPuissanceFlag,
			influxdb.HorlogeModeDegradeField:           frameData.horlogeModeDegradeFlag,
			influxdb.CommunicationEuridisField:         frameData.communicationEuridisState,
			influxdb.StatusCPLField:                    frameData.statusCPLState,
			influxdb.SynchronisationCPLField:           frameData.synchronisationCPLBool,
			influxdb.MessageField:                      frameData.messageValue,
		},
		*frameData.date)
	if err != nil {
		logger.WithError(err).Warn("Error generating power Influxdb point")
		return
	}

	*pointList = append(*pointList, stgePoint)
}

// intToString convert int to string
func intToString(number int) string {
	return strconv.FormatInt(int64(number), 10)
}

// convertTypedMapToAnyMap convert the in map to the correct type
func convertTypedMapToAnyMap[T any](inMap map[string]T) map[string]interface{} {
	outMap := make(map[string]interface{}, len(inMap))
	for k, v := range inMap {
		outMap[k] = v
	}

	return outMap
}

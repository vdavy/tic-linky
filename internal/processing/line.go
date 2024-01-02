package processing

import (
	"github.com/davecgh/go-spew/spew"
	"strings"
	"time"
)

// frameData context data struct
type frameData struct {
	date     *time.Time
	indexMap map[string]uint64
	powerMap map[string]uint64
}

// processLine process a single line od data
func (frameData *frameData) processLine(line string) {
	// preliminary checks
	// checksum
	if !validateChecksum(line) {
		logger.Warnf("Wrong checksum for %s", line)
		return
	}

	// split data into array
	splitLine := strings.Split(line, string(separatorChar))
	if len(splitLine) < 3 {
		logger.Warnf("Too short fields in split line : %d", len(splitLine))
		return
	}

	// route data
	frameData.routeLineData(splitLine)

	// finish the job at the end of frame
	if detectEndOfFrame(line) {
		frameData.processEndOfFrame()
	}
}

// routeLineData route data to the correct parser
func (frameData *frameData) routeLineData(splitLine []string) {
	switch splitLine[fieldNameIndex] {
	case dateField:
		frameData.parseDate(splitLine[dateFieldIndex])
	case eastField,
		easf01Field, easf02Field, easf03Field, easf04Field, easf05Field,
		easf06Field, easf07Field, easf08Field, easf09Field, easf10Field,
		easd01Field, easd02Field, easd03Field, easd04Field:
		parseFieldAsUint64(frameData.indexMap, splitLine[fieldNameIndex], splitLine[nonDatedFieldIndex])
	case sinstsField, urms1Field, irms1Field:
		parseFieldAsUint64(frameData.powerMap, splitLine[fieldNameIndex], splitLine[nonDatedFieldIndex])
	}
}

// processEndOfFrame send the collected data to influxdb
func (frameData *frameData) processEndOfFrame() {
	logger.Debugf("Frame data : %v", spew.Sdump(frameData))
}

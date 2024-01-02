package processing

import "strings"

const separatorChar = '\t'

// validateChecksum validate the checksum against the protocol
func validateChecksum(line string) bool {
	checksumIndex := strings.LastIndexByte(line, separatorChar)
	if checksumIndex == -1 {
		return false
	}

	var sum int
	for i := 0; i <= checksumIndex; i++ { // sum all values until last \t
		sum += int(line[i])
	}
	sum &= 0x3F // truncate 6 first bytes
	sum += 0x20 // add offset for printable char

	return byte(sum) == line[checksumIndex+1]
}

// detectEndOfFrame detect when the new frame is the beginning
func detectEndOfFrame(line string) bool {
	startTextIndex := strings.LastIndexByte(line, 0x03) // lookup ETX char
	if startTextIndex == -1 {                           // if not found
		return false
	}

	return line[startTextIndex] == 0x03 && line[startTextIndex+1] == 0x02 // should find ETX char followed by STX char
}

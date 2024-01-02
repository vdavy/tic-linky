package processing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateChecksum_True(t *testing.T) {
	line := "PJOURF+1\t00008002 0038C001 07388002 0D1A8001 0E1A8002 NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE\t]\r\x03\x02\n"
	assert.True(t, validateChecksum(line))
}

func TestValidateChecksum_False(t *testing.T) {
	line := "PJOURF+1\t00008002 0038C001 07388002 0D1A8001 0E1A8002 NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE\t[\r\x03\x02\n"
	assert.False(t, validateChecksum(line))
}

func TestValidateChecksum_True_2(t *testing.T) {
	line := "DATE\tH240101173711\t\t=\r\n"
	assert.True(t, validateChecksum(line))
}

func TestValidateChecksum_True_3(t *testing.T) {
	line := "CCASN-1\tH240101170000\t00596\t]\r\n"
	assert.True(t, validateChecksum(line))
}

func TestDetectStartOfFrame_GoodCase(t *testing.T) {
	line := "PJOURF+1\t00008002 0038C001 07388002 0D1A8001 0E1A8002 NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE\t]\r\x03\x02\n"
	assert.True(t, detectEndOfFrame(line))
}

func TestDetectStartOfFrame_AlmostGoodCase(t *testing.T) {
	line := "PJOURF+1\t00008002 0038C001 07388002 0D1A8001 0E1A8002 NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE NONUTILE\t]\r\x03\x01\n"
	assert.False(t, detectEndOfFrame(line))
}

func TestDetectStartOfFrame_Wrong(t *testing.T) {
	line := "DATE\tH240101173711\t\t=\n"
	assert.False(t, detectEndOfFrame(line))
}

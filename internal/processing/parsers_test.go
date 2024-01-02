package processing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDateParser_GoodCase(t *testing.T) {
	testFD := &frameData{}
	testFD.parseDate("H240102175338")
	assert.Equal(t, "240102175338", testFD.date.Format(dateValueFormat))
}

func TestDateParser_GoodCase_WithSpace(t *testing.T) {
	testFD := &frameData{}
	testFD.parseDate(" 240102175338")
	assert.Equal(t, "240102175338", testFD.date.Format(dateValueFormat))
}

func TestDateParser_Error(t *testing.T) {
	testFD := &frameData{}
	testFD.parseDate("xxx")
	assert.Nil(t, testFD.date)
}

func TestIndexParser_GoodCase(t *testing.T) {
	fieldMap := make(map[string]uint64)
	parseFieldAsUint64(fieldMap, eastField, "035504800")
	assert.Equal(t, uint64(35504800), fieldMap[eastField])
}

func TestIndexParser_WrongCase(t *testing.T) {
	fieldMap := make(map[string]uint64)
	parseFieldAsUint64(fieldMap, eastField, "035504ab800")
	assert.Len(t, fieldMap, 0)
	assert.Zero(t, fieldMap[eastField])
}

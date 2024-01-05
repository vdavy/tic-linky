package processing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDateParser_GoodCase(t *testing.T) {
	testFD := &frameData{}
	parseDate(&testFD.date, "H240102175338")
	assert.Equal(t, "240102175338", testFD.date.Format(dateValueFormat))
}

func TestDateParser_GoodCase_WithSpace(t *testing.T) {
	testFD := &frameData{}
	parseDate(&testFD.date, " 240102175338")
	assert.Equal(t, "240102175338", testFD.date.Format(dateValueFormat))
}

func TestDateParser_Error(t *testing.T) {
	testFD := &frameData{}
	parseDate(&testFD.date, "xxx")
	assert.Nil(t, testFD.date)
}

func TestIndexParser_GoodCase(t *testing.T) {
	fieldMap := make(map[string]int64)
	parseFieldAsInt64(fieldMap, eastField, "035504800")
	assert.Equal(t, int64(35504800), fieldMap[eastField])
}

func TestIndexParser_WrongCase(t *testing.T) {
	fieldMap := make(map[string]int64)
	parseFieldAsInt64(fieldMap, eastField, "035504ab800")
	assert.Len(t, fieldMap, 0)
	assert.Zero(t, fieldMap[eastField])
}

func TestParseDatedField(t *testing.T) {
	initCurrentFrameData()
	currentFrameData.parseDatedField([]string{"SMAXSN", "H240103130230", "05160"})
	assert.Equal(t, int64(5160), currentFrameData.datedFieldsMap["SMAXSN"].value)
	assert.Equal(t, "240103130230", currentFrameData.datedFieldsMap["SMAXSN"].date.Format(dateValueFormat))
	assert.True(t, currentFrameData.datedFieldsWriteFlagMap["SMAXSN"])
}

func TestParseDatedField_WrongHour(t *testing.T) {
	currentFrameData.datedFieldsMap = nil
	currentFrameData.datedFieldsWriteFlagMap = nil
	initCurrentFrameData()

	currentFrameData.parseDatedField([]string{"SMAXSN", "ZZZ", "05160"})
	assert.Equal(t, int64(0), currentFrameData.datedFieldsMap["SMAXSN"].value)
	assert.Nil(t, currentFrameData.datedFieldsMap["SMAXSN"].date)
	assert.False(t, currentFrameData.datedFieldsWriteFlagMap["SMAXSN"])
}

func TestParseDatedField_WrongValue(t *testing.T) {
	currentFrameData.datedFieldsMap = nil
	currentFrameData.datedFieldsWriteFlagMap = nil
	initCurrentFrameData()

	currentFrameData.parseDatedField([]string{"SMAXSN", "H240103130230", "ZZZ"})
	assert.Equal(t, int64(0), currentFrameData.datedFieldsMap["SMAXSN"].value)
	assert.Nil(t, currentFrameData.datedFieldsMap["SMAXSN"].date)
	assert.False(t, currentFrameData.datedFieldsWriteFlagMap["SMAXSN"])
}

func TestParseDatedField_SecondAdd_NoWrite(t *testing.T) {
	currentFrameData.datedFieldsMap = nil
	currentFrameData.datedFieldsWriteFlagMap = nil
	initCurrentFrameData()

	currentFrameData.parseDatedField([]string{"SMAXSN", "H240103130230", "05160"})
	assert.Equal(t, int64(5160), currentFrameData.datedFieldsMap["SMAXSN"].value)
	assert.Equal(t, "240103130230", currentFrameData.datedFieldsMap["SMAXSN"].date.Format(dateValueFormat))
	assert.True(t, currentFrameData.datedFieldsWriteFlagMap["SMAXSN"])

	currentFrameData.datedFieldsWriteFlagMap["SMAXSN"] = false
	currentFrameData.parseDatedField([]string{"SMAXSN", "H240103130230", "05161"})
	assert.Equal(t, int64(5160), currentFrameData.datedFieldsMap["SMAXSN"].value)
	assert.Equal(t, "240103130230", currentFrameData.datedFieldsMap["SMAXSN"].date.Format(dateValueFormat))
	assert.False(t, currentFrameData.datedFieldsWriteFlagMap["SMAXSN"])
}

func TestParseDatedField_SecondAdd_WithWrite(t *testing.T) {
	currentFrameData.datedFieldsMap = nil
	currentFrameData.datedFieldsWriteFlagMap = nil
	initCurrentFrameData()

	currentFrameData.parseDatedField([]string{"SMAXSN", "H240103130230", "05160"})
	assert.Equal(t, int64(5160), currentFrameData.datedFieldsMap["SMAXSN"].value)
	assert.Equal(t, "240103130230", currentFrameData.datedFieldsMap["SMAXSN"].date.Format(dateValueFormat))
	assert.True(t, currentFrameData.datedFieldsWriteFlagMap["SMAXSN"])

	currentFrameData.datedFieldsWriteFlagMap["SMAXSN"] = false
	currentFrameData.parseDatedField([]string{"SMAXSN", "H240103130231", "05161"})
	assert.Equal(t, int64(5161), currentFrameData.datedFieldsMap["SMAXSN"].value)
	assert.Equal(t, "240103130231", currentFrameData.datedFieldsMap["SMAXSN"].date.Format(dateValueFormat))
	assert.True(t, currentFrameData.datedFieldsWriteFlagMap["SMAXSN"])
}

func TestParseSTGE_Case1(t *testing.T) {
	initCurrentFrameData()
	currentFrameData.parseSTGE("003AC401")
	assert.Equal(t, 2, currentFrameData.productionIndex)
	assert.Equal(t, 4, currentFrameData.distributionIndex)

	assert.True(t, currentFrameData.contactSecOuvertFlag)
	assert.Equal(t, 0, currentFrameData.organeDeCoupureState)
	assert.False(t, currentFrameData.cacheBorneDistributeurOuvertFlag)
	assert.False(t, currentFrameData.surtensionFlag)
	assert.False(t, currentFrameData.depassementPuissanceFlag)
	assert.False(t, currentFrameData.horlogeModeDegradeFlag)
	assert.Equal(t, 3, currentFrameData.communicationEuridisState)
	assert.Equal(t, 1, currentFrameData.statusCPLState)
	assert.False(t, currentFrameData.synchronisationCPLBool)
}

func TestParseSTGE_Case2(t *testing.T) {
	initCurrentFrameData()
	currentFrameData.parseSTGE("00CB54D4")
	assert.Equal(t, 6, currentFrameData.productionIndex)
	assert.Equal(t, 2, currentFrameData.distributionIndex)

	assert.False(t, currentFrameData.contactSecOuvertFlag)
	assert.Equal(t, 2, currentFrameData.organeDeCoupureState)
	assert.True(t, currentFrameData.cacheBorneDistributeurOuvertFlag)
	assert.True(t, currentFrameData.surtensionFlag)
	assert.True(t, currentFrameData.depassementPuissanceFlag)
	assert.True(t, currentFrameData.horlogeModeDegradeFlag)
	assert.Equal(t, 1, currentFrameData.communicationEuridisState)
	assert.Equal(t, 2, currentFrameData.statusCPLState)
	assert.True(t, currentFrameData.synchronisationCPLBool)
}

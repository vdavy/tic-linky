package processing

import "time"

// frameData context data struct
type frameData struct {
	date *time.Time // global date

	// indexes data
	indexesMap        map[string]int64
	productionIndex   int
	distributionIndex int

	// power data
	powerMap map[string]int64

	// dated data
	datedFieldsMap          map[string]datedField
	datedFieldsWriteFlagMap map[string]bool

	// STGE data
	contactSecOuvertFlag             bool
	organeDeCoupureState             int
	cacheBorneDistributeurOuvertFlag bool
	surtensionFlag                   bool
	depassementPuissanceFlag         bool
	horlogeModeDegradeFlag           bool
	communicationEuridisState        int
	statusCPLState                   int
	synchronisationCPLBool           bool

	// message text
	messageValue string
}

// datedField type for dated field
type datedField struct {
	date  *time.Time
	value int64
}

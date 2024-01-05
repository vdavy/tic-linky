package processing

const (
	// production index
	productionMask        = 0x0F
	productionOffSet      = 10
	productionFinalOffSet = 1

	// distribution index
	distributionMask        = 0x03
	distributionOffSet      = 14
	distributionFinalOffSet = 1

	// organe de coupure
	organeDeCoupureMask        = 0x3
	organeDeCoupureOffset      = 1
	organeDeCoupureFinalOffset = 0

	// communication euridis
	communicationEuridisMask        = 0x3
	communicationEuridisOffset      = 19
	communicationEuridisFinalOffset = 0

	// statut CPL
	statutCPLMask        = 0x3
	statutCPLOffset      = 21
	statutCPLFinalOffset = 0
)

const (
	contactSecOffset             = 0
	cacheBorneDistributeurOffset = 4
	surtensionOffset             = 6
	depassementPuissanceOffset   = 7
	horlogeModeDegradeOffset     = 16
	synchronisationCPLOffset     = 23
)

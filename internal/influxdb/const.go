package influxdb

const (
	// IndexesMeasurement name for the indexes measurement
	IndexesMeasurement = "Indexes"
	// ProductionTag name for the production index tag
	ProductionTag = "Production"
	// DistributionTag name for the distribution index tag
	DistributionTag = "Distribution"

	// PowerMeasurement name for the power measurement
	PowerMeasurement = "Power"

	// DatedDataFieldName dated data field name used in point generation
	DatedDataFieldName = "Value"

	// ContactSecOuvertField field used in STGE point
	ContactSecOuvertField = "Contact_sec_ouvert"
	// OrganeDeCoupureField field used in STGE point
	OrganeDeCoupureField = "Organe_de_coupure"
	// CacheBorneDistributeurOuvertField field used in STGE point
	CacheBorneDistributeurOuvertField = "Cache_borne_distributeur_ouvert"
	// SurtensionField field used in STGE point
	SurtensionField = "Surtension"
	// DepassementPuissanceField field used in STGE point
	DepassementPuissanceField = "Depassement_puissance"
	// HorlogeModeDegradeField field used in STGE point
	HorlogeModeDegradeField = "Horloge_mode_degrade"
	// CommunicationEuridisField field used in STGE point
	CommunicationEuridisField = "Communication_euridis"
	// StatusCPLField field used in STGE point
	StatusCPLField = "Status_CPL"
	// SynchronisationCPLField field used in STGE point
	SynchronisationCPLField = "Synchronisation_CPL"
	// MessageField field used in STGE point (annex)
	MessageField = "Message"
)

package entity

// VetClinic represents a vet clinic
type VetClinic struct {
	ClinicName string  `json:"clinicName"`
	StateCode  string  `json:"stateCode"`
	Opening    Opening `json:"opening"`
}

// Opening represents clinic's open hours between From and To
type Opening struct {
	From string `json:"from"`
	To   string `json:"to"`
}

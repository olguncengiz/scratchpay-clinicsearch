package entity

// DentalClinic represents a dental clinic
type DentalClinic struct {
	Name         string       `json:"name"`
	StateName    string       `json:"stateName"`
	Availability Availability `json:"availability"`
}

// Availability represents clinic's open hours between From and To
type Availability struct {
	From string `json:"from"`
	To   string `json:"to"`
}

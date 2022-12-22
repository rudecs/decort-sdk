package computeci

// Main information about computeci
type ItemComputeCI struct {
	// Custom fields
	CustomFields map[string]interface{} `json:"customFields"`

	// Description
	Description string `json:"desc"`

	// List drivers
	Drivers []string `json:"drivers"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Template
	Template string `jsnn:"template"`
}

// List of computeci instances
type ListComputeCI []ItemComputeCI

package computeci

type ComputeCIRecord struct {
	CustomFields map[string]interface{} `json:"customFields"`
	Description  string                 `json:"desc"`
	Drivers      []string               `json:"drivers"`
	GUID         uint64                 `json:"guid"`
	ID           uint64                 `json:"id"`
	Name         string                 `json:"name"`
	Status       string                 `json:"status"`
	Template     string                 `jsnn:"template"`
}

type ComputeCIList []ComputeCIRecord

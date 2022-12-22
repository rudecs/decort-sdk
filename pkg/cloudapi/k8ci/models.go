package k8ci

// Detailed information about K8CI
type ItemK8CI struct {
	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Main information about K8CI
	RecordK8CI
}

// List of K8CI
type ListK8CI []ItemK8CI

// Main information about K8CI
type RecordK8CI struct {
	// Description
	Description string `json:"desc"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Version
	Version string `json:"version"`
}

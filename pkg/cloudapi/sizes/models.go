package sizes

// Main onformation about configured available flavors
type ItemSize struct {
	// Description
	Description string `json:"desc"`

	// List of disk IDs
	Disks []uint64 `json:"disks"`

	// ID
	ID uint64 `json:"id"`

	// Memory
	Memory uint64 `json:"memory"`

	// Name
	Name string `json:"name"`

	// VCPUs
	VCPUs uint64 `json:"vcpus"`
}

// List of configured available flavors
type ListSizes []ItemSize

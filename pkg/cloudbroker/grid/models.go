package grid

// Resource information
type Resources struct {
	// Current resources
	Current RecordResource `json:"Current"`

	// Reserved resources
	Reserved RecordResource `json:"Reserved"`
}

// Resource details
type RecordResource struct {
	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Disk size
	DiskSize uint64 `json:"disksize"`

	// External IPs
	ExtIPs uint64 `json:"extips"`

	// External traffic
	ExtTraffic uint64 `json:"exttraffic"`

	// Number of GPU
	GPU uint64 `json:"gpu"`

	// Number of RAM
	RAM uint64 `json:"ram"`
}

// Detailed information about grid
type RecordGrid struct {
	// Resource information
	Resources Resources `json:"Resources"`

	// Flag
	Flag string `json:"flag"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Location code
	LocationCode string `json:"locationCode"`

	// Name
	Name string `json:"name"`
}

// List Grids
type ListGrids []RecordGrid

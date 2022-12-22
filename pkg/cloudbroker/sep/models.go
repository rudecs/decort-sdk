package sep

// Total resource information
type Total struct {
	// Capacity limit
	CapacityLimit uint64 `json:"capacity_limit"`

	// Disk count
	DiskCount uint64 `json:"disk_count"`

	// Disk usage
	DiskUsage uint64 `json:"disk_usage"`

	// Snapshot count
	SnapshotCount uint64 `json:"snapshot_count"`

	// Snapshot usage
	SnapshotUsage uint64 `json:"snapshot_usage"`

	// Usage
	Usage uint64 `json:"usage"`

	// Usage limit
	UsageLimit uint64 `json:"usage_limit"`
}

// Main information about consumption
type RecordConsumption struct {
	// By pool
	ByPool map[string]interface{} `json:"byPool"`

	// Total resource information
	Total Total `json:"total"`

	// Type
	Type string `json:"type"`
}

// Main information about URI
type ItemURI struct {
	// IP
	IP string `json:"ip"`

	// Port
	Port uint64 `json:"port"`
}

// List URIs
type ListURIs []ItemURI

// Detailed information about SEP pool
type RecordPool struct {
	// List access account IDs
	AccessAccountIDs []uint64 `json:"accessAccountIds"`

	// List access resource group IDs
	AccessResGroupIDs []uint64 `json:"accessResGroupIds"`

	// Name
	Name string `json:"name"`

	//  Page cache ratio
	PageCacheRatio uint64 `json:"pagecache_ratio"`

	// Reference ID
	ReferenceID string `json:"referenceId"`

	// List types
	Types []string `json:"types"`

	// List URIs
	URIs ListURIs `json:"uris"`
}

// SEP config
type SEPConfig map[string]interface{}

// Detailed information about SEP
type RecordSEP struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Config
	Config SEPConfig `json:"config"`

	// Consumed by
	ConsumedBy []uint64 `json:"consumedBy"`

	// Description
	Description string `json:"desc"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Object status
	ObjStatus string `json:"objStatus"`

	// Provided by
	ProvidedBy []uint64 `json:"providedBy"`

	// Shared with
	SharedWith []uint64 `json:"sharedWith"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`
}

// List SEPs
type ListSEP []RecordSEP

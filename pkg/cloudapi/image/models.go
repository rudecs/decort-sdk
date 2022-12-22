package image

// Main information about image
type ItemImage struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Architecture
	Architecture string `json:"architecture"`

	// Boot type
	BootType string `json:"bootType"`

	// Bootable
	Bootable bool `json:"bootable"`

	// CDROM
	CDROM bool `json:"cdrom"`

	// Description
	Description string `json:"desc"`

	// List drivers
	Drivers []string `json:"drivers"`

	// HotResize
	HotResize bool `json:"hotResize"`

	// ID
	ID uint64 `json:"id"`

	// Link to
	LinkTo uint64 `json:"linkTo"`

	// Name
	Name string `json:"name"`

	// Pool
	Pool string `json:"pool"`

	// SepID
	SepID uint64 `json:"sepId"`

	// Size
	Size uint64 `json:"size"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`

	// Username
	Username string `json:"username"`

	// Virtual
	Virtual bool `json:"virtual"`
}

// List of information about images
type ListImages []ItemImage

// History
type History struct {
	// GUID
	GUID string `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Timestamp
	Timestamp uint64 `json:"timestamp"`
}

// Detailed information about image
type RecordImage struct {
	// UNCPathj
	UNCPath string `json:"UNCPath"`

	// CKey
	CKey string `json:"_ckey"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Access Control List
	ACL interface{} `json:"acl"`

	// Architecture
	Architecture string `json:"architecture"`

	// Boot type
	BootType string `json:"bootType"`

	// Bootable
	Bootable bool `json:"bootable"`

	// ComputeCI ID
	ComputeCIID uint64 `json:"computeciId"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// List of drivers
	Drivers []string `json:"drivers"`

	// Enabled
	Enabled bool `json:"enabled"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// History
	History []History `json:"history"`

	// HotResize
	HotResize bool `json:"hotResize"`

	// ID
	ID uint64 `json:"id"`

	// Last modified
	LastModified uint64 `json:"lastModified"`

	// Link to
	LinkTo uint64 `json:"linkTo"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Password
	Password string `json:"password"`

	// Pool
	Pool string `json:"pool"`

	// ProviderName
	ProviderName string `json:"provider_name"`

	// Purge attempts
	PurgeAttempts uint64 `json:"purgeAttempts"`

	// Resource ID
	ResID string `json:"resId"`

	// RescueCD
	RescueCD bool `json:"rescuecd"`

	// SepID
	SepID uint64 `json:"sepId"`

	// SharedWith list
	SharedWith []uint64 `json:"sharedWith"`

	// Size
	Size uint64 `json:"size"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`

	// Username
	Username string `json:"username"`

	// Version
	Version string `json:"version"`
}

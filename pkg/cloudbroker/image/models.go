package image

// Detailed information about image
type RecordImage struct {
	// UNC path
	UNCPath string `json:"UNCPath"`

	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Access Control List
	ACL ListACL `json:"acl"`

	// Architecture
	Architecture string `json:"architecture"`

	// Boot type
	BootType string `json:"bootType"`

	// Bootable
	Bootable bool `json:"bootable"`

	// Compute CI ID
	ComputeCIID uint64 `json:"computeciId"`

	// Deleted  time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// Drivers
	Drivers []string `json:"drivers"`

	// Enabled
	Enabled bool `json:"enabled"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// List history
	History ListHistory `json:"history"`

	// Hot resize
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

	// Provider name
	ProviderName string `json:"provider_name"`

	// Purge attempts
	PurgeAttempts uint64 `json:"purgeAttempts"`

	// Reference ID
	ReferenceID string `json:"referenceId"`

	// Resource ID
	ResID string `json:"resId"`

	// Resource name
	ResName string `json:"resName"`

	// Rescue CD
	RescueCD bool `json:"rescuecd"`

	// SEP ID
	SEPID uint64 `json:"sepId"`

	// List shared with
	SharedWith []uint64 `json:"sharedWith"`

	// Size
	Size uint64 `json:"size"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`

	// URL
	URL string `json:"url"`

	// Username
	Username string `json:"username"`

	// Version
	Version string `json:"version"`

	// Virtual
	Virtual bool `json:"virtual"`
}

// List images
type ListImages []RecordImage

// Access Control List
type ACL struct {
	// Explicit
	Explicit bool `json:"explicit"`

	// GUID
	GUID string `json:"guid"`

	// Right
	Right string `json:"right"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`

	// User group ID
	UserGroupID string `json:"userGroupId"`
}

// List ACL
type ListACL []ACL

// History information
type History struct {
	// GUID
	GUID string `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Timestamp
	Timestamp uint64 `json:"timestamp"`
}

// List history
type ListHistory []History

// List stacks
type ListStacks []struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// API URL
	APIURL string `json:"apiUrl"`

	// API key
	APIKey string `json:"apikey"`

	// App ID
	AppID string `json:"appId"`

	// Description
	Description string `json:"desc"`

	// Drivers
	Drivers []string `json:"drivers"`

	// Eco
	Eco interface{} `json:"eco"`

	// Error
	Error uint64 `json:"error"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// List image IDs
	Images []uint64 `json:"images"`

	// Login
	Login string `json:"login"`

	// Name
	Name string `json:"name"`

	// Password
	Password string `json:"passwd"`

	// Reference ID
	ReferenceID string `json:"referenceId"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`
}

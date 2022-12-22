package locations

// Main information about locations
type ItemLocation struct {
	// Grid ID
	GID uint64 `json:"gid"`

	// ID
	ID uint64 `json:"id"`

	// GUID
	GUID uint64 `json:"guid"`

	// Location code
	LocationCode string `json:"locationCode"`

	// Name
	Name string `json:"name"`

	// Flag
	Flag string `json:"flag"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// CKey
	CKey string `json:"_ckey"`
}

// List of locations
type ListLocations []ItemLocation

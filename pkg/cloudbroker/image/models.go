package image

type ImageRecord struct {
	UNCPath       string        `json:"UNCPath"`
	CKey          string        `json:"_ckey"`
	Meta          []interface{} `json:"_meta"`
	AccountID     uint64        `json:"accountId"`
	ACL           []ACL         `json:"acl"`
	Architecture  string        `json:"architecture"`
	BootType      string        `json:"bootType"`
	Bootable      bool          `json:"bootable"`
	ComputeCIID   uint64        `json:"computeciId"`
	DeletedTime   uint64        `json:"deletedTime"`
	Desc          string        `json:"desc"`
	Drivers       []string      `json:"drivers"`
	Enabled       bool          `json:"enabled"`
	Gid           uint64        `json:"gid"`
	GUID          uint64        `json:"guid"`
	History       []History     `json:"history"`
	HotResize     bool          `json:"hotResize"`
	ID            uint64        `json:"id"`
	LastModified  uint64        `json:"lastModified"`
	LinkTo        uint64        `json:"linkTo"`
	Milestones    uint64        `json:"milestones"`
	Name          string        `json:"name"`
	Password      string        `json:"password"`
	Pool          string        `json:"pool"`
	ProviderName  string        `json:"provider_name"`
	PurgeAttempts uint64        `json:"purgeAttempts"`
	ReferenceID   string        `json:"referenceId"`
	ResID         string        `json:"resId"`
	ResName       string        `json:"resName"`
	RescueCD      bool          `json:"rescuecd"`
	SepID         uint64        `json:"sepId"`
	SharedWith    []uint64      `json:"sharedWith"`
	Size          uint64        `json:"size"`
	Status        string        `json:"status"`
	TechStatus    string        `json:"techStatus"`
	Type          string        `json:"type"`
	URL           string        `json:"url"`
	Username      string        `json:"username"`
	Version       string        `json:"version"`
	Virtual       bool          `json:"virtual"`
}

type ListImages []ImageRecord

type ACL struct {
	Explicit    bool   `json:"explicit"`
	GUID        string `json:"guid"`
	Right       string `json:"right"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	UserGroupID string `json:"userGroupId"`
}
type History struct {
	GUID      string `json:"guid"`
	ID        uint64 `json:"id"`
	Timestamp uint64 `json:"timestamp"`
}

type ListStacks []struct {
	CKey        string        `json:"_ckey"`
	Meta        []interface{} `json:"_meta"`
	APIURL      string        `json:"apiUrl"`
	APIKey      string        `json:"apikey"`
	AppID       string        `json:"appId"`
	Desc        string        `json:"desc"`
	Drivers     []string      `json:"drivers"`
	Eco         interface{}   `json:"eco"`
	Error       uint64        `json:"error"`
	GID         uint64        `json:"gid"`
	GUID        uint64        `json:"guid"`
	ID          uint64        `json:"id"`
	Images      []uint64      `json:"images"`
	Login       string        `json:"login"`
	Name        string        `json:"name"`
	Passwd      string        `json:"passwd"`
	ReferenceID string        `json:"referenceId"`
	Status      string        `json:"status"`
	Type        string        `json:"type"`
}

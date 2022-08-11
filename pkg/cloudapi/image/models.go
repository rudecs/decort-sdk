package image

type ImageRecord struct {
	AccountId    uint64   `json:"accountId"`
	Architecture string   `json:"architecture"`
	BootType     string   `json:"bootType"`
	Bootable     bool     `json:"bootable"`
	CDROM        bool     `json:"cdrom"`
	Description  string   `json:"desc"`
	Drivers      []string `json:"drivers"`
	HotResize    bool     `json:"hotResize"`
	ID           uint64   `json:"id"`
	LinkTo       uint64   `json:"linkTo"`
	Name         string   `json:"name"`
	Pool         string   `json:"pool"`
	SepId        uint64   `json:"sepId"`
	Size         uint64   `json:"size"`
	Status       string   `json:"status"`
	Type         string   `json:"type"`
	Username     string   `json:"username"`
	Virtual      bool     `json:"virtual"`
}

type ImageList []ImageRecord

type History struct {
	GUID      string `json:"guid"`
	ID        uint64 `json:"id"`
	Timestamp uint64 `json:"timestamp"`
}

type ImageExtend struct {
	UNCPath       string      `json:"UNCPath"`
	CKey          string      `json:"_ckey"`
	AccountId     uint64      `json:"accountId"`
	Acl           interface{} `json:"acl"`
	Architecture  string      `json:"architecture"`
	BootType      string      `json:"bootType"`
	Bootable      bool        `json:"bootable"`
	ComputeCiId   uint64      `json:"computeciId"`
	DeletedTime   uint64      `json:"deletedTime"`
	Description   string      `json:"desc"`
	Drivers       []string    `json:"drivers"`
	Enabled       bool        `json:"enabled"`
	GID           uint64      `json:"gid"`
	GUID          uint64      `json:"guid"`
	History       []History   `json:"history"`
	HotResize     bool        `json:"hotResize"`
	ID            uint64      `json:"id"`
	LastModified  uint64      `json:"lastModified"`
	LinkTo        uint64      `json:"linkTo"`
	Milestones    uint64      `json:"milestones"`
	Name          string      `json:"name"`
	Password      string      `json:"password"`
	Pool          string      `json:"pool"`
	ProviderName  string      `json:"provider_name"`
	PurgeAttempts int         `json:"purgeAttempts"`
	ResId         string      `json:"resId"`
	RescueCD      bool        `json:"rescuecd"`
	SepId         uint64      `json:"sepId"`
	SharedWith    []int       `json:"sharedWith"`
	Size          uint64      `json:"size"`
	Status        string      `json:"status"`
	TechStatus    string      `json:"techStatus"`
	Type          string      `json:"type"`
	Username      string      `json:"username"`
	Version       string      `json:"version"`
}
package flipgroup

type FlipGroupRecord struct {
	DefaultGW string `json:"defaultGW"`
	ID        uint64 `json:"id"`
	IP        string `json:"ip"`
	Name      string `json:"name"`
	Netmask   uint64 `json:"netmask"`
}

type FlipGroupItem struct {
	AccountID   uint64   `json:"accountId"`
	AccountName string   `json:"accountName"`
	ClientIDs   []uint64 `json:"clientIds"`
	ClientType  string   `json:"clientType"`
	ConnID      uint64   `json:"connId"`
	ConnType    string   `json:"connType"`
	CreatedBy   string   `json:"createdBy"`
	CreatedTime uint64   `json:"createdTime"`
	DefaultGW   string   `json:"defaultGW"`
	DeletedBy   string   `json:"deletedBy"`
	DeletedTime uint64   `json:"deletedTime"`
	Description string   `json:"desc"`
	GID         uint64   `json:"gid"`
	GUID        uint64   `json:"guid"`
	ID          uint64   `json:"id"`
	IP          string   `json:"ip"`
	Milestones  uint64   `json:"milestones"`
	Name        string   `json:"name"`
	NetID       uint64   `json:"netId"`
	NetType     string   `json:"netType"`
	Network     string   `json:"network"`
	RGID        uint64   `json:"rgId"`
	RGName      string   `json:"rgName"`
	Status      string   `json:"status"`
	UpdatedBy   string   `json:"updatedBy"`
	UpdatedTime uint64   `json:"updatedTime"`
}

type FlipGroupList []FlipGroupItem

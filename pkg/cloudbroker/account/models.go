package account

type AccountAudit struct {
	Call         string  `json:"call"`
	ResponseTime float64 `json:"responsetime"`
	StatusCode   uint64  `json:"statuscode"`
	Timestamp    float64 `json:"timestamp"`
	User         string  `json:"user"`
}

type AccountAuditsList []AccountAudit

type Resources struct {
	Current  Current  `json:"Current"`
	Reserved Reserved `json:"Reserved"`
}

type Current struct {
	CPU        uint64 `json:"cpu"`
	DiskSize   uint64 `json:"disksize"`
	ExtIPs     uint64 `json:"extips"`
	ExtTraffic uint64 `json:"exttraffic"`
	GPU        uint64 `json:"gpu"`
	RAM        uint64 `json:"ram"`
}

type Reserved struct {
	CPU        uint64 `json:"cpu"`
	DiskSize   uint64 `json:"disksize"`
	ExtIPs     uint64 `json:"extips"`
	ExtTraffic uint64 `json:"exttraffic"`
	GPU        uint64 `json:"gpu"`
	RAM        uint64 `json:"ram"`
}

type ACL struct {
	Explicit    bool   `json:"explicit"`
	GUID        string `json:"guid"`
	Right       string `json:"right"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	UserGroupID string `json:"userGroupId"`
}

type ResourceLimits struct {
	CuC      float64 `json:"CU_C"`
	CuD      float64 `json:"CU_D"`
	CuI      float64 `json:"CU_I"`
	CuM      float64 `json:"CU_M"`
	CuNP     float64 `json:"CU_NP"`
	GPUUnits float64 `json:"gpu_units"`
}

type InfoResponse struct {
	DCLocation       string         `json:"DCLocation"`
	CKey             string         `json:"_ckey"`
	ACL              []ACL          `json:"acl"`
	Company          string         `json:"company"`
	CompanyURL       string         `json:"companyurl"`
	CreatedBy        string         `json:"createdBy"`
	CreatedTime      uint64         `json:"createdTime"`
	DeactivationTime float64        `json:"deactivationTime"`
	DeletedBy        string         `json:"deletedBy"`
	DeletedTime      uint64         `json:"deletedTime"`
	DisplayName      string         `json:"displayname"`
	GUID             uint64         `json:"guid"`
	ID               uint64         `json:"id"`
	Name             string         `json:"name"`
	ResourceLimits   ResourceLimits `json:"resourceLimits"`
	SendAccessEmails bool           `json:"sendAccessEmails"`
	Status           string         `json:"status"`
	UpdatedTime      uint64         `json:"updatedTime"`
	Version          uint64         `json:"version"`
	VINS             []uint64       `json:"vins"`
}
type GetResponse struct {
	Resources Resources `json:"Resources"`
	InfoResponse
}

type ListInfoResponse []struct {
	Meta []interface{} `json:"_meta"`
	InfoResponse
}

type ListComputes []Compute
type Compute struct {
	AccountID      uint64 `json:"accountId"`
	AccountName    string `json:"accountName"`
	CPUs           uint64 `json:"cpus"`
	CreatedBy      string `json:"createdBy"`
	CreatedTime    uint64 `json:"createdTime"`
	DeletedBy      string `json:"deletedBy"`
	DeletedTime    uint64 `json:"deletedTime"`
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	RAM            uint64 `json:"ram"`
	Registered     bool   `json:"registered"`
	RgID           uint64 `json:"rgId"`
	RgName         string `json:"rgName"`
	Status         string `json:"status"`
	TechStatus     string `json:"techStatus"`
	TotalDisksSize uint64 `json:"totalDisksSize"`
	UpdatedBy      string `json:"updatedBy"`
	UpdatedTime    uint64 `json:"updatedTime"`
	UserManaged    bool   `json:"userManaged"`
	VINSConnected  uint64 `json:"vinsConnected"`
}

type ListDisks []Disk

type Disk struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Pool    string `json:"pool"`
	SepID   uint64 `json:"sepId"`
	SizeMax uint64 `json:"sizeMax"`
	Type    string `json:"type"`
}

type ListFlipGroups []FlipGroup

type FlipGroup struct {
	AccountID   uint64 `json:"accountId"`
	ClientType  string `json:"clientType"`
	ConnType    string `json:"connType"`
	CreatedBy   string `json:"createdBy"`
	CreatedTime uint64 `json:"createdTime"`
	DefaultGW   string `json:"defaultGW"`
	DeletedBy   string `json:"deletedBy"`
	DeletedTime uint64 `json:"deletedTime"`
	Desc        string `json:"desc"`
	Gid         uint64 `json:"gid"`
	GUID        uint64 `json:"guid"`
	ID          uint64 `json:"id"`
	IP          string `json:"ip"`
	Milestones  uint64 `json:"milestones"`
	Name        string `json:"name"`
	NetID       uint64 `json:"netId"`
	NetType     string `json:"netType"`
	Netmask     uint64 `json:"netmask"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime uint64 `json:"updatedTime"`
}

type Computes struct {
	Started uint64 `json:"Started"`
	Stopped uint64 `json:"Stopped"`
}

type Consumed struct {
	CPU        uint64 `json:"cpu"`
	DiskSize   uint64 `json:"disksize"`
	ExtIPs     uint64 `json:"extips"`
	ExtTraffic uint64 `json:"exttraffic"`
	GPU        uint64 `json:"gpu"`
	RAM        uint64 `json:"ram"`
}

type Limits struct {
	CPU        int64 `json:"cpu"`
	DiskSize   int64 `json:"disksize"`
	ExtIPs     int64 `json:"extips"`
	ExtTraffic int64 `json:"exttraffic"`
	GPU        int64 `json:"gpu"`
	RAM        int64 `json:"ram"`
}

type RGResuorces struct {
	Consumed Consumed `json:"Consumed"`
	Limits   Limits   `json:"Limits"`
	Reserved Reserved `json:"Reserved"`
}

type RG struct {
	Computes    Computes    `json:"Computes"`
	Resources   RGResuorces `json:"Resources"`
	CreatedBy   string      `json:"createdBy"`
	CreatedTime uint64      `json:"createdTime"`
	DeletedBy   string      `json:"deletedBy"`
	DeletedTime uint64      `json:"deletedTime"`
	ID          uint64      `json:"id"`
	Milestones  uint64      `json:"milestones"`
	Name        string      `json:"name"`
	Status      string      `json:"status"`
	UpdatedBy   string      `json:"updatedBy"`
	UpdatedTime uint64      `json:"updatedTime"`
	VINSes      uint64      `json:"vinses"`
}

type ListRG []RG

type VINS struct {
	AccountID   uint64 `json:"accountId"`
	AccountName string `json:"accountName"`
	Computes    uint64 `json:"computes"`
	CreatedBy   string `json:"createdBy"`
	CreatedTime uint64 `json:"createdTime"`
	DeletedBy   string `json:"deletedBy"`
	DeletedTime uint64 `json:"deletedTime"`
	ExternalIP  string `json:"externalIP"`
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Network     string `json:"network"`
	PriVnfDevID uint64 `json:"priVnfDevId"`
	RgID        uint64 `json:"rgId"`
	RgName      string `json:"rgName"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime uint64 `json:"updatedTime"`
}

type ListVINS []VINS

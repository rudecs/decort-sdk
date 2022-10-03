package account

type AccountACLRecord struct {
	IsExplicit   bool   `json:"explicit"`
	GUID         string `json:"guid"`
	Rights       string `json:"right"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	UgroupID     string `json:"userGroupId"`
	CanBeDeleted bool   `json:"canBeDeleted"`
}

type ResourceLimits struct {
	CUC      float64 `json:"CU_C"`
	CUD      float64 `json:"CU_D"`
	CUI      float64 `json:"CU_I"`
	CUM      float64 `json:"CU_M"`
	CUNP     float64 `json:"CU_NP"`
	GPUUnits float64 `json:"gpu_units"`
}

type AccountRecord struct {
	DCLocation        string             `json:"DCLocation"`
	CKey              string             `jspn:"_ckey"`
	Meta              []interface{}      `json:"_meta"`
	ACL               []AccountACLRecord `json:"acl"`
	Company           string             `json:"company"`
	CompanyURL        string             `json:"companyurl"`
	CreatedBy         string             `jspn:"createdBy"`
	CreatedTime       uint64             `json:"createdTime"`
	DeactiovationTime float64            `json:"deactivationTime"`
	DeletedBy         string             `json:"deletedBy"`
	DeletedTime       uint64             `json:"deletedTime"`
	DisplayName       string             `json:"displayname"`
	GUID              uint64             `json:"guid"`
	ID                uint64             `json:"id"`
	Name              string             `json:"name"`
	ResourceLimits    ResourceLimits     `json:"resourceLimits"`
	SendAccessEmails  bool               `json:"sendAccessEmails"`
	ServiceAccount    bool               `json:"serviceAccount"`
	Status            string             `json:"status"`
	UpdatedTime       uint64             `json:"updatedTime"`
	Version           uint64             `json:"version"`
	VINS              []uint64           `json:"vins"`
}

type AccountList []AccountRecord

type AccountCloudApi struct {
	ACL         []AccountACLRecord `json:"acl"`
	CreatedTime uint64             `json:"createdTime"`
	DeletedTime uint64             `json:"deletedTime"`
	ID          uint64             `json:"id"`
	Name        string             `json:"name"`
	Status      string             `json:"status"`
	UpdatedTime uint64             `json:"updatedTime"`
}

type AccountCloudApiList []AccountCloudApi

type Resource struct {
	CPU        int64 `json:"cpu"`
	DiskSize   int64 `json:"disksize"`
	ExtIPs     int64 `json:"extips"`
	ExtTraffic int64 `json:"exttraffic"`
	GPU        int64 `json:"gpu"`
	RAM        int64 `json:"ram"`
}

type Resources struct {
	Current  Resource `json:"Current"`
	Reserved Resource `json:"Reserved"`
}

type Computes struct {
	Started uint64 `json:"started"`
	Stopped uint64 `json:"stopped"`
}

type Machines struct {
	Running uint64 `json:"running"`
	Halted  uint64 `json:"halted"`
}

type AccountWithResources struct {
	Account
	Resources Resources `json:"Resources"`
	Computes  Computes  `json:"computes"`
	Machines  Machines  `json:"machines"`
	VINSes    uint64    `json:"vinses"`
}

type AccountCompute struct {
	AccountID      uint64 `json:"accountId"`
	AccountName    string `json:"accountName"`
	CPUs           uint64 `json:"cpus"`
	CreatedBy      string `json:"createdBy"`
	CreatedTime    uint64 `json:"createdTime"`
	DeletedBy      string `json:"deletedBy"`
	DeletedTime    uint64 `json:"deletedTime"`
	ComputeID      uint64 `json:"id"`
	ComputeName    string `json:"name"`
	RAM            uint64 `json:"ram"`
	Registered     bool   `json:"registered"`
	RGID           uint64 `json:"rgId"`
	RGName         string `json:"rgName"`
	Status         string `json:"status"`
	TechStatus     string `json:"techStatus"`
	TotalDisksSize uint64 `json:"totalDisksSize"`
	UpdatedBy      string `json:"updatedBy"`
	UpdatedTime    uint64 `json:"updatedTime"`
	UserManaged    bool   `json:"userManaged"`
	VINSConnected  uint64 `json:"vinsConnected"`
}

type AccountComputesList []AccountCompute

type AccountDisk struct {
	ID      uint64 `json:"id"`
	Name    string `json:"name"`
	Pool    string `json:"pool"`
	SepID   uint64 `json:"sepId"`
	SizeMax uint64 `json:"sizeMax"`
	Type    string `json:"type"`
}

type AccountDisksList []AccountDisk

type AccountVIN struct {
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
	RGID        uint64 `json:"rgId"`
	RGName      string `json:"rgName"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime uint64 `json:"updatedTime"`
}

type AccountVINSList []AccountVIN

type AccountAudit struct {
	Call         string  `json:"call"`
	ResponseTime float64 `json:"responsetime"`
	StatusCode   uint64  `json:"statuscode"`
	Timestamp    float64 `json:"timestamp"`
	User         string  `json:"user"`
}

type AccountAuditsList []AccountAudit

type AccountRGComputes struct {
	Started uint64 `json:"Started"`
	Stopped uint64 `json:"Stopped"`
}

type AccountRGResources struct {
	Consumed Resource `json:"Consumed"`
	Limits   Resource `json:"Limits"`
	Reserved Resource `json:"Reserved"`
}

type AccountRG struct {
	Computes    AccountRGComputes  `json:"Computes"`
	Resources   AccountRGResources `json:"Resources"`
	CreatedBy   string             `json:"createdBy"`
	CreatedTime uint64             `json:"createdTime"`
	DeletedBy   string             `json:"deletedBy"`
	DeletedTime uint64             `json:"deletedTime"`
	RGID        uint64             `json:"id"`
	Milestones  uint64             `json:"milestones"`
	RGName      string             `json:"name"`
	Status      string             `json:"status"`
	UpdatedBy   string             `json:"updatedBy"`
	UpdatedTime uint64             `json:"updatedTime"`
	VINSes      uint64             `json:"vinses"`
}

type AccountRGList []AccountRG

type AccountTemplate struct {
	UNCPath     string `json:"UNCPath"`
	AccountID   uint64 `json:"accountId"`
	Description string `json:"desc"`
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Public      bool   `json:"public"`
	Size        uint64 `json:"size"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Username    string `json:"username"`
}

type AccountTemplatesList []AccountTemplate

type AccountFlipGroup struct {
	AccountID   uint64 `json:"accountId"`
	ClientType  string `json:"clientType"`
	ConnType    string `json:"connType"`
	CreatedBy   string `json:"createdBy"`
	CreatedTime uint64 `json:"createdTime"`
	DefaultGW   string `json:"defaultGW"`
	DeletedBy   string `json:"deletedBy"`
	DeletedTime uint64 `json:"deletedTime"`
	Description string `json:"desc"`
	GID         uint64 `json:"gid"`
	GUID        uint64 `json:"guid"`
	ID          uint64 `json:"id"`
	IP          string `json:"ip"`
	Milestones  uint64 `json:"milestones"`
	Name        string `json:"name"`
	NetID       uint64 `json:"netId"`
	NetType     string `json:"netType"`
	NetMask     uint64 `json:"netmask"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime uint64 `json:"updatedTime"`
}

type AccountFlipGroupsList []AccountFlipGroup

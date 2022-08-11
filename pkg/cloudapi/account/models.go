package account

type AccountAclRecord struct {
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
	GpuUnits float64 `json:"gpu_units"`
}

type AccountRecord struct {
	DCLocation        string             `json:"DCLocation"`
	CKey              string             `jspn:"_ckey"`
	Meta              []interface{}      `json:"_meta"`
	Acl               []AccountAclRecord `json:"acl"`
	Company           string             `json:"company"`
	CompanyUrl        string             `json:"companyurl"`
	CreatedBy         string             `jspn:"createdBy"`
	CreatedTime       int                `json:"createdTime"`
	DeactiovationTime float64            `json:"deactivationTime"`
	DeletedBy         string             `json:"deletedBy"`
	DeletedTime       int                `json:"deletedTime"`
	DisplayName       string             `json:"displayname"`
	GUID              int                `json:"guid"`
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	ResourceLimits    ResourceLimits     `json:"resourceLimits"`
	SendAccessEmails  bool               `json:"sendAccessEmails"`
	ServiceAccount    bool               `json:"serviceAccount"`
	Status            string             `json:"status"`
	UpdatedTime       int                `json:"updatedTime"`
	Version           int                `json:"version"`
	Vins              []int              `json:"vins"`
}

type AccountList []AccountRecord

type AccountCloudApi struct {
	Acl         []AccountAclRecord `json:"acl"`
	CreatedTime int                `json:"createdTime"`
	DeletedTime int                `json:"deletedTime"`
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Status      string             `json:"status"`
	UpdatedTime int                `json:"updatedTime"`
}

type AccountCloudApiList []AccountCloudApi

type Resource struct {
	CPU        int `json:"cpu"`
	Disksize   int `json:"disksize"`
	Extips     int `json:"extips"`
	Exttraffic int `json:"exttraffic"`
	GPU        int `json:"gpu"`
	RAM        int `json:"ram"`
}

type Resources struct {
	Current  Resource `json:"Current"`
	Reserved Resource `json:"Reserved"`
}

type Computes struct {
	Started int `json:"started"`
	Stopped int `json:"stopped"`
}

type Machines struct {
	Running int `json:"running"`
	Halted  int `json:"halted"`
}

type AccountWithResources struct {
	Account
	Resources Resources `json:"Resources"`
	Computes  Computes  `json:"computes"`
	Machines  Machines  `json:"machines"`
	Vinses    int       `json:"vinses"`
}

type AccountCompute struct {
	AccountId      int    `json:"accountId"`
	AccountName    string `json:"accountName"`
	CPUs           int    `json:"cpus"`
	CreatedBy      string `json:"createdBy"`
	CreatedTime    int    `json:"createdTime"`
	DeletedBy      string `json:"deletedBy"`
	DeletedTime    int    `json:"deletedTime"`
	ComputeId      int    `json:"id"`
	ComputeName    string `json:"name"`
	RAM            int    `json:"ram"`
	Registered     bool   `json:"registered"`
	RGId           int    `json:"rgId"`
	RGName         string `json:"rgName"`
	Status         string `json:"status"`
	TechStatus     string `json:"techStatus"`
	TotalDisksSize int    `json:"totalDisksSize"`
	UpdatedBy      string `json:"updatedBy"`
	UpdatedTime    int    `json:"updatedTime"`
	UserManaged    bool   `json:"userManaged"`
	VinsConnected  int    `json:"vinsConnected"`
}

type AccountComputesList []AccountCompute

type AccountDisk struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Pool    string `json:"pool"`
	SepId   int    `json:"sepId"`
	SizeMax int    `json:"sizeMax"`
	Type    string `json:"type"`
}

type AccountDisksList []AccountDisk

type AccountVin struct {
	AccountId   int    `json:"accountId"`
	AccountName string `json:"accountName"`
	Computes    int    `json:"computes"`
	CreatedBy   string `json:"createdBy"`
	CreatedTime int    `json:"createdTime"`
	DeletedBy   string `json:"deletedBy"`
	DeletedTime int    `json:"deletedTime"`
	ExternalIP  string `json:"externalIP"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Network     string `json:"network"`
	PriVnfDevId int    `json:"priVnfDevId"`
	RGId        int    `json:"rgId"`
	RGName      string `json:"rgName"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime int    `json:"updatedTime"`
}

type AccountVinsList []AccountVin

type AccountAudit struct {
	Call         string  `json:"call"`
	ResponseTime float64 `json:"responsetime"`
	StatusCode   int     `json:"statuscode"`
	Timestamp    float64 `json:"timestamp"`
	User         string  `json:"user"`
}

type AccountAuditsList []AccountAudit

type AccountRGComputes struct {
	Started int `json:"Started"`
	Stopped int `json:"Stopped"`
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
	CreatedTime int                `json:"createdTime"`
	DeletedBy   string             `json:"deletedBy"`
	DeletedTime int                `json:"deletedTime"`
	RGID        int                `json:"id"`
	Milestones  int                `json:"milestones"`
	RGName      string             `json:"name"`
	Status      string             `json:"status"`
	UpdatedBy   string             `json:"updatedBy"`
	UpdatedTime int                `json:"updatedTime"`
	Vinses      int                `json:"vinses"`
}

type AccountRGList []AccountRG

type AccountTemplate struct {
	UNCPath     string `json:"UNCPath"`
	AccountId   int    `json:"accountId"`
	Description string `json:"desc"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Public      bool   `json:"public"`
	Size        int    `json:"size"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Username    string `json:"username"`
}

type AccountTemplatesList []AccountTemplate

type AccountFlipGroup struct {
	AccountId   int    `json:"accountId"`
	ClientType  string `json:"clientType"`
	ConnType    string `json:"connType"`
	CreatedBy   string `json:"createdBy"`
	CreatedTime int    `json:"createdTime"`
	DefaultGW   string `json:"defaultGW"`
	DeletedBy   string `json:"deletedBy"`
	DeletedTime int    `json:"deletedTime"`
	Description string `json:"desc"`
	GID         int    `json:"gid"`
	GUID        int    `json:"guid"`
	ID          int    `json:"id"`
	IP          string `json:"ip"`
	Milestones  int    `json:"milestones"`
	Name        string `json:"name"`
	NetID       int    `json:"netId"`
	NetType     string `json:"netType"`
	NetMask     int    `json:"netmask"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime int    `json:"updatedTime"`
}

type AccountFlipGroupsList []AccountFlipGroup
package account

// Access Control List
type RecordACL struct {
	// Whether access is explicitly specified
	IsExplicit bool `json:"explicit"`

	// GUID
	GUID string `json:"guid"`

	// Access rights
	Rights string `json:"right"`

	// Status
	Status string `json:"status"`

	// Account Type
	Type string `json:"type"`

	// Account owner ID
	UgroupID string `json:"userGroupId"`

	// Is it possible to remove
	CanBeDeleted bool `json:"canBeDeleted"`
}

// Resource limits
type ResourceLimits struct {
	// Number of cores
	CUC float64 `json:"CU_C"`

	// Disk size, GB
	CUD float64 `json:"CU_D"`

	// Number of public IP addresses
	CUI float64 `json:"CU_I"`

	// RAM size, MB
	CUM float64 `json:"CU_M"`

	// Traffic volume, GB
	CUNP float64 `json:"CU_NP"`

	// Number of graphics cores
	GPUUnits float64 `json:"gpu_units"`
}

// Main information of account
type InfoAccount struct {
	// Segment
	DCLocation string `json:"DCLocation"`

	// Key
	CKey string `jspn:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Access Control List
	ACL []RecordACL `json:"acl"`

	// Company
	Company string `json:"company"`

	// Company URL
	CompanyURL string `json:"companyurl"`

	// Created by
	CreatedBy string `jspn:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deactiovation time
	DeactiovationTime float64 `json:"deactivationTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Display name
	DisplayName string `json:"displayname"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Resource Limits
	ResourceLimits ResourceLimits `json:"resourceLimits"`

	// If true send emails when a user is granted access to resources
	SendAccessEmails bool `json:"sendAccessEmails"`

	// Service Account
	ServiceAccount bool `json:"serviceAccount"`

	// Status
	Status string `json:"status"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// Version
	Version uint64 `json:"version"`

	// List VINS in account
	VINS []uint64 `json:"vins"`
}

// Main information in one of if the list of accounts
type ItemAccount struct {
	// Access Control List
	ACL []RecordACL `json:"acl"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// List of accounts
type ListAccounts []ItemAccount

// Resources used
type Resource struct {
	// Number of cores
	CPU int64 `json:"cpu"`

	// Disk size
	DiskSize int64 `json:"disksize"`

	// Number of External IPs
	ExtIPs int64 `json:"extips"`

	// External traffic
	ExtTraffic int64 `json:"exttraffic"`

	// Number of grafic cores
	GPU int64 `json:"gpu"`

	// Number of RAM
	RAM int64 `json:"ram"`
}

// Information about resources
type Resources struct {
	// Current information about resources
	Current Resource `json:"Current"`

	// Reserved information about resources
	Reserved Resource `json:"Reserved"`
}

// Information about computes
type Computes struct {
	// Number of started computes
	Started uint64 `json:"started"`

	// Number of stopped computes
	Stopped uint64 `json:"stopped"`
}

// Information about machines
type Machines struct {
	// Number of running machines
	Running uint64 `json:"running"`

	// Number of halted machines
	Halted uint64 `json:"halted"`
}

// Сomplete information about account
type RecordAccount struct {
	// Main information about account
	InfoAccount

	// Resources
	Resources Resources `json:"Resources"`

	// Computes
	Computes Computes `json:"computes"`

	// Machines
	Machines Machines `json:"machines"`

	// Number of VINSes
	VINSes uint64 `json:"vinses"`
}

// Main information about compute
type ItemCompute struct {
	// ID an account
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Number of CPU
	CPUs uint64 `json:"cpus"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// ID compute
	ComputeID uint64 `json:"id"`

	// Compute name
	ComputeName string `json:"name"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// Registered or not
	Registered bool `json:"registered"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group Name
	RGName string `json:"rgName"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Total disks size
	TotalDisksSize uint64 `json:"totalDisksSize"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// User controlled or not
	UserManaged bool `json:"userManaged"`

	// Number of connected VINS
	VINSConnected uint64 `json:"vinsConnected"`
}

// List of computes
type ListComputes []ItemCompute

// Main information about disk
type ItemDisk struct {
	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Pool
	Pool string `json:"pool"`

	// ID SEP
	SEPID uint64 `json:"sepId"`

	// Max size
	SizeMax uint64 `json:"sizeMax"`

	// Disk type
	Type string `json:"type"`
}

// List of disks
type ListDisks []ItemDisk

// Main information about VINS
type ItemVINS struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Name of account
	AccountName string `json:"accountName"`

	// Number of computes
	Computes uint64 `json:"computes"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// External IP
	ExternalIP string `json:"externalIP"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Network
	Network string `json:"network"`

	// NNFDev ID
	PriVNFDevID uint64 `json:"priVnfDevId"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// Status
	Status string `json:"status"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// List of VINS
type ListVINS []ItemVINS

// Main info about audit
type ItemAudit struct {
	// Call
	Call string `json:"call"`

	// Response time
	ResponseTime float64 `json:"responsetime"`

	// Status code
	StatusCode uint64 `json:"statuscode"`

	// Timestamp
	Timestamp float64 `json:"timestamp"`

	// User
	User string `json:"user"`
}

// List of audits
type ListAudits []ItemAudit

// Information compute in resource group
type RGComputes struct {
	// Number of started computes
	Started uint64 `json:"started"`

	// Number of stopped computes
	Stopped uint64 `json:"stopped"`
}

// Resources of Resource group
type RGResources struct {
	// Consumed
	Consumed Resource `json:"Consumed"`

	// Limits
	Limits Resource `json:"Limits"`

	// Reserved
	Reserved Resource `json:"Reserved"`
}

// Main information about resource group
type ItemRG struct {
	// Computes
	Computes RGComputes `json:"Computes"`

	// Resources
	Resources RGResources `json:"Resources"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Resource group ID
	RGID uint64 `json:"id"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Resource group name
	RGName string `json:"name"`

	// Status
	Status string `json:"status"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// Number of VINS
	VINSes uint64 `json:"vinses"`
}

// List of Resource groups
type ListRG []ItemRG

// Main information about template
type ItemTemplate struct {
	// UNCPath
	UNCPath string `json:"UNCPath"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Description
	Description string `json:"desc"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Public or not
	Public bool `json:"public"`

	// Size
	Size uint64 `json:"size"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`

	// Username
	Username string `json:"username"`
}

// List of templates
type ListTemplates []ItemTemplate

// Main information about FLIPGroup
type ItemFLIPGroup struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Client type
	ClientType string `json:"clientType"`

	// Connection type
	ConnType string `json:"connType"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Default GW
	DefaultGW string `json:"defaultGW"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// IP
	IP string `json:"ip"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Network ID
	NetID uint64 `json:"netId"`

	// Network type
	NetType string `json:"netType"`

	// Network mask
	NetMask uint64 `json:"netmask"`

	// Status
	Status string `json:"status"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// List of FLIPGroups
type ListFLIPGroups []ItemFLIPGroup

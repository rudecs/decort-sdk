package account

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

type RecordResources struct {
	// Current information about resources
	Current Resource `json:"Current"`

	// Reserved information about resources
	Reserved Resource `json:"Reserved"`
}

type Resource struct {
	// Number of cores
	CPU int64 `json:"cpu"`

	// Disk size
	DiskSize float64 `json:"disksize"`

	// Disk size max
	DiskSizeMax uint64 `json:"disksizemax"`

	// Number of External IPs
	ExtIPs int64 `json:"extips"`

	// External traffic
	ExtTraffic int64 `json:"exttraffic"`

	// Number of grafic cores
	GPU int64 `json:"gpu"`

	// Number of RAM
	RAM int64 `json:"ram"`

	// SEPs
	SEPs map[string]map[string]DiskUsage `json:"seps"`
}

// Disk usage
type DiskUsage struct {
	// Disk size
	DiskSize float64 `json:"disksize"`

	// Disk size max
	DiskSizeMax uint64 `json:"disksizemax"`
}

// Access Control List
type ACL struct {
	// Whether access is explicitly specified
	Explicit bool `json:"explicit"`

	// GUID
	GUID string `json:"guid"`

	// Access rights
	Right string `json:"right"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`

	// User group ID
	UserGroupID string `json:"userGroupId"`
}

// Resource limits
type ResourceLimits struct {
	// CuC
	CuC float64 `json:"CU_C"`

	// CuD
	CuD float64 `json:"CU_D"`

	// CuI
	CuI float64 `json:"CU_I"`

	// CuM
	CuM float64 `json:"CU_M"`

	// CuNP
	CuNP float64 `json:"CU_NP"`

	// GPUUnits
	GPUUnits float64 `json:"gpu_units"`
}

// Main information about account
type InfoAccount struct {
	// DCLocation
	DCLocation string `json:"DCLocation"`

	// CKey
	CKey string `json:"_ckey"`

	// Access Control List
	ACL []ACL `json:"acl"`

	// Company
	Company string `json:"company"`

	// Company URL
	CompanyURL string `json:"companyurl"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deactivation time
	DeactivationTime float64 `json:"deactivationTime"`

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

	// Resource limits
	ResourceLimits ResourceLimits `json:"resourceLimits"`

	// Resource types
	ResourceTypes []string `json:"resourceTypes"`

	// Send access emails
	SendAccessEmails bool `json:"sendAccessEmails"`

	// Status
	Status string `json:"status"`

	// UniqPools
	UniqPools []string `json:"uniqPools"`

	// UpdatedTime
	UpdatedTime uint64 `json:"updatedTime"`

	// Version
	Version uint64 `json:"version"`

	// List of VINS IDs
	VINS []uint64 `json:"vins"`
}

// Deatailed information about account
type RecordAccount struct {
	// Resources
	Resources RecordResources `json:"Resources"`

	// Main information about account
	InfoAccount
}

// More information about account
type ItemAccount struct {
	// Meta
	Meta []interface{} `json:"_meta"`

	// Main information about account
	InfoAccount
}

// List of accounts
type ListAccounts []ItemAccount

// List of computes
type ListComputes []ItemCompute

// Main information about compute
type ItemCompute struct {
	// Account ID
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

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// Registered
	Registered bool `json:"registered"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RgName string `json:"rgName"`

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

	// User managed
	UserManaged bool `json:"userManaged"`

	// VINS Connected
	VINSConnected uint64 `json:"vinsConnected"`
}

// List of disks
type ListDisks []ItemDisk

// Main information about disks
type ItemDisk struct {
	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Pool
	Pool string `json:"pool"`

	// SepID
	SepID uint64 `json:"sepId"`

	// Shareable
	Shareable bool `json:"shareable"`

	// Size max
	SizeMax uint64 `json:"sizeMax"`

	// Type
	Type string `json:"type"`
}

// List of FLIPGroups
type ListFLIPGroups []ItemFLIPGroup

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
	Netmask uint64 `json:"netmask"`

	// Status
	Status string `json:"status"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// Computes info
type Computes struct {
	// Started
	Started uint64 `json:"Started"`

	// Stopped
	Stopped uint64 `json:"Stopped"`
}

// Consumed
type Consumed struct {
	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Disk size
	DiskSize uint64 `json:"disksize"`

	// Disk size max
	DiskSizeMax int64 `json:"disksizemax"`

	// External IPs
	ExtIPs uint64 `json:"extips"`

	// External traffic
	ExtTraffic uint64 `json:"exttraffic"`

	// Number of GPU
	GPU uint64 `json:"gpu"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// SEPs
	SEPs map[string]map[string]DiskUsage `json:"seps"`
}

// Limits
type Limits struct {
	// Number of CPU
	CPU int64 `json:"cpu"`

	// Disk size
	DiskSize int64 `json:"disksize"`

	// Disk size max
	DiskSizeMax int64 `json:"disksizemax"`

	// External IPs
	ExtIPs int64 `json:"extips"`

	// External traffic
	ExtTraffic int64 `json:"exttraffic"`

	// Number of GPU
	GPU int64 `json:"gpu"`

	// Number of RAM
	RAM int64 `json:"ram"`

	// SEPs number
	SEPs uint64 `json:"seps"`
}

// Resources of resource group
type RGResuorces struct {
	//  Consumed
	Consumed Consumed `json:"Consumed"`

	// Limits
	Limits Limits `json:"Limits"`

	// Reserved
	Reserved Resource `json:"Reserved"`
}

// Main information about Resource group
type ItemRG struct {
	// Compute
	Computes Computes `json:"Computes"`

	// Resources of resource group
	Resources RGResuorces `json:"Resources"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// ID
	ID uint64 `json:"id"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// Number of VINSes
	VINSes uint64 `json:"vinses"`
}

// List of resource groups
type ListRG []ItemRG

// Main information about VINS
type ItemVINS struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Computes
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

	// PriVNFDevID
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

// List of VINSes
type ListVINS []ItemVINS

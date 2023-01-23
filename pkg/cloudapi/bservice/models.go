package bservice

// Detailed info about BasicService
type RecordBasicService struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Base domain
	BaseDomain string `json:"baseDomain"`

	// List Computes
	Computes ListComputes `json:"computes"`

	// Number of cores
	CPUTotal uint64 `json:"cpuTotal"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Amount of disk space used, GB
	DiskTotal uint64 `json:"diskTotal"`

	// Grid ID
	GID uint64 `json:"gid"`

	// List of Service Compute Group IDs
	Groups []uint64 `json:"groups"`

	// List of compute groups by name
	GroupsName []string `json:"groupsName"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Parent service ID
	ParentSrvID uint64 `json:"parentSrvId"`

	// Parent service type
	ParentSrvType string `json:"parentSrvType"`

	// Total amount of RAM, MB
	RAMTotal uint64 `json:"ramTotal"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// List of snapshots
	Snapshots ListSnapshots `json:"snapshots"`

	// SSH key for connection
	SSHKey string `json:"sshKey"`

	// Username for SSH connection
	SSHUser string `json:"sshUser"`

	// status
	Status string `json:"status"`

	// TechStatus
	TechStatus string `json:"techStatus"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// Whether user controlled
	UserManaged bool `json:"userManaged"`
}

// Main information about Compute
type ItemCompute struct {
	// Compute group ID
	CompGroupID uint64 `json:"compgroupId"`

	// Compute group name
	CompGroupName string `json:"compgroupName"`

	// compute group role
	CompGroupRole string `json:"compgroupRole"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`
}

// List of Computes
type ListComputes []ItemCompute

// Main information about Snapshot
type ItemSnapshot struct {
	// GUID
	GUID string `json:"guid"`

	// Label
	Label string `json:"label"`

	// Timestamp
	Timestamp uint64 `json:"timestamp"`

	// Valid or not
	Valid bool `json:"valid"`
}

// List of Snapshots
type ListSnapshots []ItemSnapshot

// Main information about Group
type RecordGroup struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account Name
	AccountName string `json:"accountName"`

	// List of Computes
	Computes ListGroupComputes `json:"computes"`

	// Consistency or not
	Consistency bool `json:"consistency"`

	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Amount of disk
	Disk uint64 `json:"disk"`

	// Driver
	Driver string `json:"driver"`

	// list of External Network IDs
	ExtNets []uint64 `json:"extnets"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Image ID
	ImageID uint64 `json:"imageId"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// List of Parent IDs
	Parents []uint64 `json:"parents"`

	// Pool name
	PoolName string `json:"poolName"`

	// Number of RAM, MB
	RAM uint64 `json:"ram"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// Role
	Role string `json:"role"`

	// SEPID
	SEPID uint64 `json:"sepId"`

	// Sequence number
	SeqNo uint64 `json:"seqNo"`

	// Service ID
	ServiceID uint64 `json:"serviceId"`

	// Status
	Status string `json:"status"`

	// TechStatus
	TechStatus string `json:"techStatus"`

	// Timeout Start
	TimeoutStart uint64 `json:"timeoutStart"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// List of VINS IDs
	VINSes []uint64 `json:"vinses"`
}

// Main information about Group Compute
type ItemGroupCompute struct {
	// ID
	ID uint64 `json:"id"`

	// IP Addresses
	IPAddresses []string `json:"ipAddresses"`

	// Name
	Name string `json:"name"`

	// List of information about OS Users
	OSUsers ListOSUsers `json:"osUsers"`
}

// List of Group Computes
type ListGroupComputes []ItemGroupCompute

// Main information about OS User
type ItemOSUser struct {
	// Login
	Login string `json:"login"`

	// Password
	Password string `json:"password"`
}

// List of information about OS Users
type ListOSUsers []ItemOSUser

// Main information about BasicService
type ItemBasicService struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Base domain
	BaseDomain string `json:"baseDomain"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Grid ID
	GID uint64 `json:"gid"`

	// List of group IDs
	Groups []uint64 `json:"groups"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Parent service ID
	ParentSrvID uint64 `json:"parentSrvId"`

	// Parent service type
	ParentSrvType string `json:"parentSrvType"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// SSH user
	SSHUser string `json:"sshUser"`

	// Status
	Status string `json:"status"`

	// TechStatus
	TechStatus string `json:"techStatus"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// User Managed or not
	UserManaged bool `json:"userManaged"`
}

// List of BasicServices
type ListBasicServices []ItemBasicService

package k8s

// Deteiled information
type ItemDetailedInfo struct {
	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Status
	TechStatus string `json:"techStatus"`
}

// List detailed information
type ListDetailedInfo []ItemDetailedInfo

// Detailed information about K8S group
type RecordK8SGroup struct {
	// List annotations
	Annotations []string `json:"annotations"`

	// Number of CPU
	CPU uint64 `json:"cpu"`

	// List detailed information
	DetailedInfo ListDetailedInfo `json:"detailedInfo"`

	// Disk
	Disk uint64 `json:"disk"`

	// GUID
	GUID string `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// List labels
	Labels []string `json:"labels"`

	// Name
	Name string `json:"name"`

	// Number
	Num uint64 `json:"num"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// List taints
	Taints []string `json:"taints"`
}

// List K8S groups
type ListK8SGroup []RecordK8SGroup

// Detailed information about K8S
type RecordK8S struct {
	// Access Control List
	ACL RecordACLGroup `json:"ACL"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Basic service ID
	BServiceID uint64 `json:"bserviceId"`

	// CI ID
	CIID uint64 `json:"ciId"`

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

	// K8CI name
	K8CIName string `json:"k8ciName"`

	// Detailed information about K8S groups
	K8SGroups RecordK8SGroups `json:"k8sGroups"`

	// Load balancer ID
	LBID uint64 `json:"lbId"`

	// Name
	Name string `json:"name"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// Detailed info about K8S groups
type RecordK8SGroups struct {
	// Master group
	Masters MasterGroup `json:"masters"`

	// Worker group
	Workers ListK8SGroup `json:"workers"`
}

// Detailed information about master group
type MasterGroup struct {
	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Detailed information
	DetailedInfo ListDetailedInfo `json:"detailedInfo"`

	// Disk
	Disk uint64 `json:"disk"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Number
	Num uint64 `json:"num"`

	// Number of RAM
	RAM uint64 `json:"ram"`
}

// Detailed information of access control
type RecordACLGroup struct {
	// Account ACL
	AccountACL ListACL `json:"accountAcl"`

	// K8S ACL
	K8SACL ListACL `json:"k8sAcl"`

	// RG ACL
	RGACL ListACL `json:"rgAcl"`
}

// Access Control List
type ACL struct {
	// Explicit
	Explicit bool `json:"explicit"`

	// GUID
	GUID string `json:"guid"`

	// Right
	Right string `json:"right"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`

	// User group ID
	UserGroupID string `json:"userGroupId"`
}

// List ACL
type ListACL []ACL

// Main information about K8S
type ItemK8S struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Access Control List
	ACL []interface{} `json:"acl"`

	// Basic service ID
	BServiceID uint64 `json:"bserviceId"`

	// CI ID
	CIID uint64 `json:"ciId"`

	// Config
	Config interface{} `json:"config"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// External network ID
	ExtNetID uint64 `json:"extnetId"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Load balancer ID
	LBID uint64 `json:"lbId"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// Service account
	ServiceAccount ServiceAccount `json:"serviceAccount"`

	// SSH key
	SSHKey string `json:"sshKey"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// VINS ID
	VINSID uint64 `json:"vinsId"`

	// List workers
	WorkersGroup ListK8SGroup `json:"workersGroups"`
}

// Service account
type ServiceAccount struct {
	// GUID
	GUID string `json:"guid"`

	// Password
	Password string `json:"password"`

	// Username
	Username string `json:"username"`
}

// List K8S
type ListK8S []ItemK8S

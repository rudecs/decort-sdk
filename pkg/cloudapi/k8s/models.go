package k8s

// Main information about kubernetes cluster
type ItemK8SGroup struct {
	// List of Annotations
	Annotations []string `json:"annotations"`

	// Number of CPU
	CPU uint64 `json:"cpu"`

	// List detailed information
	DetailedInfo ListDetailedInfo `json:"detailedInfo"`

	// Disk ID
	Disk uint64 `json:"disk"`

	// GUID
	GUID string `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// List of Labels
	Labels []string `json:"labels"`

	// Name
	Name string `json:"name"`

	// Num
	Num uint64 `json:"num"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// List of taints
	Taints []string `json:"taints"`
}

// List kubernetes cluster groups
type ListK8SGroups []ItemK8SGroup

// Detailed information
type ItemDetailedInfo struct {
	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`
}

// List detailed information
type ListDetailedInfo []ItemDetailedInfo

// Deteal information about kubernetes cluster
type RecordK8S struct {
	// Access Control List
	ACL RecordACL `json:"ACL"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Basic Service ID
	BServiceID uint64 `json:"bserviceId"`

	// CIID
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

	// Kubernetes cluster groups information
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

// Detailed information about kubernetes cluster groups
type RecordK8SGroups struct {

	// Master information
	Masters MasterGroup `json:"masters"`

	// Worker information
	Workers ListK8SGroups `json:"workers"`
}

// Master group information
type MasterGroup struct {
	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Detailed information
	DetailedInfo ListDetailedInfo `json:"detailedInfo"`

	// Disk ID
	Disk uint64 `json:"disk"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Num
	Num uint64 `json:"num"`

	// Number of RAM
	RAM uint64 `json:"ram"`
}

// Access Control List
type RecordACL struct {
	// Account ACL
	AccountACL ListACL `json:"accountAcl"`

	// K8S ACL
	K8SACL ListACL `json:"k8sAcl"`

	// RG ACL
	RGACL ListACL `json:"rgAcl"`
}

// Main information of Access Control List
type ItemACL struct {
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

// List of ACL
type ListACL []ItemACL

// Main information about kubernetes cluster
type ItemK8SCluster struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Access Control List
	ACL []interface{} `json:"acl"`

	// Basic Service ID
	BServiceID uint64 `json:"bserviceId"`

	// CIID
	CIID uint64 `json:"ciId"`

	// Config
	Config interface{} `json:"config"`

	// Create by
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

	// Information about service account
	ServiceAccount RecordServiceAccount `json:"serviceAccount"`

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

	// List workers group
	WorkersGroup ListK8SGroups `json:"workersGroups"`
}

// Information about service account
type RecordServiceAccount struct {
	// GUID
	GUID string `json:"guid"`

	// Password
	Password string `json:"password"`

	// Username
	Username string `json:"username"`
}

// List of kubernetes clusters
type ListK8SClusters []ItemK8SCluster

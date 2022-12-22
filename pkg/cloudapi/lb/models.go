package lb

// Detailed information about load balancer
type RecordLB struct {
	//HAMode
	HAMode bool `json:"HAmode"`

	// Access Control List
	ACL interface{} `json:"acl"`

	// List of load balancer backends
	Backends ListBackends `json:"backends"`

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

	// DPAPIUser
	DPAPIUser string `json:"dpApiUser"`

	// External network ID
	ExtNetID uint64 `json:"extnetId"`

	// List of load balancer frontends
	Frontends ListFrontends `json:"frontends"`

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

	// Primary node
	PrimaryNode RecordNode `json:"primaryNode"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// Secondary node
	SecondaryNode RecordNode `json:"secondaryNode"`

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
}

// Main information about load balancer
type ItemLoadBalancer struct {
	// DPAPIPassword
	DPAPIPassword string `json:"dpApiPassword"`
	// Detailed information about load balancer
	RecordLB
}

// List of load balancers
type ListLB []ItemLoadBalancer

// Main information about backend
type ItemBackend struct {
	// Algorithm
	Algorithm string `json:"algorithm"`

	// GUID
	GUID string `json:"guid"`

	// Name
	Name string `json:"name"`

	// Server settings
	ServerDefaultSettings RecordServerSettings `json:"serverDefaultSettings"`

	// List of servers
	Servers ListServers `json:"servers"`
}

// List of backends
type ListBackends []ItemBackend

// Server settings
type RecordServerSettings struct {
	// Inter
	Inter uint64 `json:"inter"`

	// GUID
	GUID string `json:"guid"`

	// DownInter
	DownInter uint64 `json:"downinter"`

	// Rise
	Rise uint64 `json:"rise"`

	// Fall
	Fall uint64 `json:"fall"`

	// SlowStart
	SlowStart uint64 `json:"slowstart"`

	// Max connections
	MaxConn uint64 `json:"maxconn"`

	// Max queue
	MaxQueue uint64 `json:"maxqueue"`

	// Weight
	Weight uint64 `json:"weight"`
}

// Main information about server
type ItemServer struct {

	// Address
	Address string `json:"address"`

	// Check
	Check string `json:"check"`

	// GUID
	GUID string `json:"guid"`

	// Name
	Name string `json:"name"`

	// Port
	Port uint64 `json:"port"`

	// Server settings
	ServerSettings RecordServerSettings `json:"serverSettings"`
}

// List of servers
type ListServers []ItemServer

// Main information about node
type RecordNode struct {
	// Backend IP
	BackendIP string `json:"backendIp"`

	// Compute ID
	ComputeID uint64 `json:"computeId"`

	// Frontend IP
	FrontendIP string `json:"frontendIp"`

	// GUID
	GUID string `json:"guid"`

	// MGMTIP
	MGMTIP string `json:"mgmtIp"`

	// Network ID
	NetworkID uint64 `json:"networkId"`
}

// Main information about frontend
type ItemFrontend struct {
	// Backend
	Backend string `json:"backend"`

	// List of bindings
	Bindings ListBindings `json:"bindings"`

	// GUID
	GUID string `json:"guid"`

	// Name
	Name string `json:"name"`
}

// List of frontends
type ListFrontends []ItemFrontend

// Main information about bindings
type ItemBinding struct {
	// Address
	Address string `json:"address"`

	// GUID
	GUID string `json:"guid"`

	// Name
	Name string `json:"name"`

	// Port
	Port uint64 `json:"port"`
}

// List of bindings
type ListBindings []ItemBinding

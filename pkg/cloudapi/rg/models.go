package rg

type ResourceGroup struct {
	AccountID        uint64         `json:"accountId"`
	AccountName      string         `json:"accountName"`
	ACL              []ACL          `json:"acl"`
	CreatedBy        string         `json:"createdBy"`
	CreatedTime      uint64         `json:"createdTime"`
	DefNetID         uint64         `json:"def_net_id"`
	DefNetType       string         `json:"def_net_type"`
	DeletedBy        string         `json:"deletedBy"`
	DeletedTime      uint64         `json:"deletedTime"`
	Desc             string         `json:"desc"`
	Dirty            bool           `url:"dirty"`
	GID              uint64         `json:"gid"`
	GUID             uint64         `json:"guid"`
	ID               uint64         `json:"id"`
	LockStatus       string         `json:"lockStatus"`
	Milestones       int            `json:"milestones"`
	Name             string         `json:"name"`
	RegisterComputes bool           `json:"registerComputes"`
	ResourceLimits   ResourceLimits `json:"resourceLimits"`
	Secret           string         `json:"secret"`
	Status           string         `json:"status"`
	UpdatedBy        string         `json:"updatedBy"`
	UpdatedTime      uint64         `json:"updatedTime"`
	VINS             []uint64       `json:"vins"`
	Computes         []uint64       `json:"vms"`
}

type ResourceGroupList []ResourceGroup

type ACL struct {
	Explicit    bool   `json:"explicit"`
	GUID        string `json:"guid"`
	Right       string `json:"right"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	UserGroupID string `json:"userGroupId"`
}

type ResourceLimits struct {
	CUC      float64 `json:"CU_C"`
	CUD      float64 `json:"CU_D"`
	CUI      float64 `json:"CU_I"`
	CUM      float64 `json:"CU_M"`
	CUNP     float64 `json:"CU_NP"`
	GPUUnits float64 `json:"gpu_units"`
}

type AffinityGroupCompute struct {
	ComputeID             uint64   `json:"computeId"`
	OtherNode             []uint64 `json:"otherNode"`
	OtherNodeIndirect     []uint64 `json:"otherNodeIndirect"`
	OtherNodeIndirectSoft []uint64 `json:"otherNodeIndirectSoft"`
	OtherNodeSoft         []uint64 `json:"otherNodeSoft"`
	SameNode              []uint64 `json:"sameNode"`
	SameNodeSoft          []uint64 `json:"sameNodeSoft"`
}

type AffinityGroupComputeList []AffinityGroupCompute

type Audit struct {
	Call         string  `json:"call"`
	ResponseTime float64 `json:"responsetime"`
	StatusCode   uint64  `json:"statuscode"`
	Timestamp    float64 `json:"timestamp"`
	User         string  `json:"user"`
}

type AuditList []Audit

type Compute struct {
	AccountID     uint64 `json:"accountId"`
	AccountName   string `json:"accountName"`
	AffinityLabel string `json:"affinityLabel"`
	// TODO put actual type here
	AffinityRules  []any  `json:"affinityRules"`
	AffinityWeight uint64 `json:"affinityWeight"`
	// TODO put actual type here
	AntiAffinityRules []any  `json:"antiAffinityRules"`
	CPUs              uint64 `json:"cpus"`
	CreatedBy         string `json:"createdBy"`
	CreatedTime       uint64 `json:"createdTime"`
	DeletedBy         string `json:"deletedBy"`
	DeletedTime       uint64 `json:"deletedTime"`
	ID                uint64 `json:"id"`
	Name              string `json:"name"`
	RAM               uint64 `json:"ram"`
	Registered        bool   `json:"registered"`
	RGID              uint64 `json:"rgId"`
	RGName            string `json:"rgName"`
	Status            string `json:"status"`
	TechStatus        string `json:"techStatus"`
	TotalDisksSize    uint64 `json:"totalDisksSize"`
	UpdatedBy         string `json:"updatedBy"`
	UpdatedTime       uint64 `json:"updatedTime"`
	UserManaged       bool   `json:"userManaged"`
	VINSConnected     uint64 `json:"vinsConnected"`
}

type ComputeList []Compute

type LoadBalancer struct {
	HAMode        bool        `json:"HAmode"`
	ACL           interface{} `json:"acl"`
	Backends      []Backend   `json:"backends"`
	CreatedBy     string      `json:"createdBy"`
	CreatedTime   uint64      `json:"createdTime"`
	DeletedBy     string      `json:"deletedBy"`
	DeletedTime   uint64      `json:"deletedTime"`
	Description   string      `json:"desc"`
	DPAPIUser     string      `json:"dpApiUser"`
	ExtnetId      uint64      `json:"extnetId"`
	Frontends     []Frontend  `json:"frontends"`
	GID           uint64      `json:"gid"`
	GUID          uint64      `json:"guid"`
	ID            uint64      `json:"id"`
	ImageId       uint64      `json:"imageId"`
	Milestones    uint64      `json:"milestones"`
	Name          string      `json:"name"`
	PrimaryNode   Node        `json:"primaryNode"`
	RGID          uint64      `json:"rgId"`
	RGName        string      `json:"rgName"`
	SecondaryNode Node        `json:"secondaryNode"`
	Status        string      `json:"status"`
	TechStatus    string      `json:"techStatus"`
	UpdatedBy     string      `json:"updatedBy"`
	UpdatedTime   uint64      `json:"updatedTime"`
	VinsId        uint64      `json:"vinsId"`
}

type LoadBalancerDetailed struct {
	DPAPIPassword string `json:"dpApiPassword"`
	LoadBalancer
}

type Backend struct {
	Algorithm             string         `json:"algorithm"`
	GUID                  string         `json:"guid"`
	Name                  string         `json:"name"`
	ServerDefaultSettings ServerSettings `json:"serverDefaultSettings"`
	Servers               []Server       `json:"servers"`
}

type LBList []LoadBalancerDetailed

type ServerSettings struct {
	Inter     uint64 `json:"inter"`
	GUID      string `json:"guid"`
	DownInter uint64 `json:"downinter"`
	Rise      uint   `json:"rise"`
	Fall      uint   `json:"fall"`
	SlowStart uint64 `json:"slowstart"`
	MaxConn   uint   `json:"maxconn"`
	MaxQueue  uint   `json:"maxqueue"`
	Weight    uint   `json:"weight"`
}

type Server struct {
	Address        string         `json:"address"`
	Check          string         `json:"check"`
	GUID           string         `json:"guid"`
	Name           string         `json:"name"`
	Port           uint           `json:"port"`
	ServerSettings ServerSettings `json:"serverSettings"`
}

type Node struct {
	BackendIp  string `json:"backendIp"`
	ComputeId  uint64 `json:"computeId"`
	FrontendIp string `json:"frontendIp"`
	GUID       string `json:"guid"`
	MGMTIp     string `json:"mgmtIp"`
	NetworkId  uint64 `json:"networkId"`
}

type Frontend struct {
	Backend  string    `json:"backend"`
	Bindings []Binding `json:"bindings"`
	GUID     string    `json:"guid"`
	Name     string    `json:"name"`
}

type Binding struct {
	Address string `json:"address"`
	GUID    string `json:"guid"`
	Name    string `json:"name"`
	Port    uint   `json:"port"`
}

type PortForward struct {
	PublicPortEnd   uint16 `json:"Public Port End"`
	PublicPortStart uint16 `json:"Public Port Start"`
	VMID            uint64 `json:"VM ID"`
	VMIP            string `json:"VM IP"`
	VMName          string `json:"VM Name"`
	VMPort          uint16 `json:"VM Port"`
	VINSID          uint64 `json:"ViNS ID"`
	VINSName        string `json:"ViNS Name"`
}

type PortForwardList []PortForward

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
	PriVNFDevID uint64 `json:"priVnfDevId"`
	RGID        uint64 `json:"rgId"`
	RGName      string `json:"rgName"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime uint64 `json:"updatedTime"`
}

type VINSList []VINS

type ResourceUsage struct {
	CPU          uint64 `json:"cpu"`
	DiskSize     uint64 `json:"disksize"`
	ExtIPs       uint64 `json:"extips"`
	ExtraTraffic uint64 `json:"exttraffic"`
	GPU          uint64 `json:"gpu"`
	RAM          uint64 `json:"ram"`
}

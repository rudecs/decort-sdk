package rg

type Audit struct {
	Call         string  `json:"call"`
	ResponseTime float64 `json:"responsetime"`
	StatusCode   uint64  `json:"statuscode"`
	Timestamp    float64 `json:"timestamp"`
	User         string  `json:"user"`
}
type AuditList []Audit

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

type Resources struct {
	Current  Current  `json:"Current"`
	Reserved Reserved `json:"Reserved"`
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
	CuNp     float64 `json:"CU_NP"`
	GPUUnits float64 `json:"gpu_units"`
}

type ResourceGroup struct {
	Resources Resources `json:"Resources"`
	InfoResponse
}

type InfoResponse struct {
	AccountID        uint64         `json:"accountId"`
	AccountName      string         `json:"accountName"`
	ACL              []ACL          `json:"acl"`
	CreatedBy        string         `json:"createdBy"`
	CreatedTime      uint64         `json:"createdTime"`
	DefNetID         int64          `json:"def_net_id"`
	DefNetType       string         `json:"def_net_type"`
	DeletedBy        string         `json:"deletedBy"`
	DeletedTime      uint64         `json:"deletedTime"`
	Desc             string         `json:"desc"`
	GID              uint64         `json:"gid"`
	GUID             uint64         `json:"guid"`
	ID               uint64         `json:"id"`
	LockStatus       string         `json:"lockStatus"`
	Milestones       uint64         `json:"milestones"`
	Name             string         `json:"name"`
	RegisterComputes bool           `json:"registerComputes"`
	ResourceLimits   ResourceLimits `json:"resourceLimits"`
	Secret           string         `json:"secret"`
	Status           string         `json:"status"`
	UpdatedBy        string         `json:"updatedBy"`
	UpdatedTime      uint64         `json:"updatedTime"`
	VINS             []interface{}  `json:"vins"`
	VMs              []interface{}  `json:"vms"`
}

type List []InfoResponse

type ListDeleted []InfoResponse

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

type AffinityRules struct {
	GUID     string `json:"guid"`
	Key      string `json:"key"`
	Mode     string `json:"mode"`
	Policy   string `json:"policy"`
	Topology string `json:"topology"`
	Value    string `json:"value"`
}

type Compute struct {
	AccountID         uint64          `json:"accountId"`
	AccountName       string          `json:"accountName"`
	AffinityLabel     string          `json:"affinityLabel"`
	AffinityRules     []AffinityRules `json:"affinityRules"`
	AffinityWeight    uint64          `json:"affinityWeight"`
	AntiAffinityRules []interface{}   `json:"antiAffinityRules"`
	CPUs              uint64          `json:"cpus"`
	CreatedBy         string          `json:"createdBy"`
	CreatedTime       uint64          `json:"createdTime"`
	DeletedBy         string          `json:"deletedBy"`
	DeletedTime       uint64          `json:"deletedTime"`
	ID                uint64          `json:"id"`
	Name              string          `json:"name"`
	RAM               uint64          `json:"ram"`
	Registered        bool            `json:"registered"`
	RGID              uint64          `json:"rgId"`
	RGName            string          `json:"rgName"`
	Status            string          `json:"status"`
	TechStatus        string          `json:"techStatus"`
	TotalDisksSize    uint64          `json:"totalDisksSize"`
	UpdatedBy         string          `json:"updatedBy"`
	UpdatedTime       uint64          `json:"updatedTime"`
	UserManaged       bool            `json:"userManaged"`
	VINSConnected     uint64          `json:"vinsConnected"`
}

type ListComputes []Compute

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

type ListVINS []VINS

type PFW struct {
	PublicPortEnd   uint64 `json:"Public Port End"`
	PublicPortStart uint64 `json:"Public Port Start"`
	VMID            uint64 `json:"VM ID"`
	VMIP            string `json:"VM IP"`
	VMName          string `json:"VM Name"`
	VMPort          uint64 `json:"VM Port"`
	VINSID          uint64 `json:"ViNS ID"`
	VINSName        string `json:"ViNS Name"`
}

type ListPFW []PFW

type ServerSettings struct {
	DownInter uint64 `json:"downinter"`
	Fall      uint64 `json:"fall"`
	GUID      string `json:"guid"`
	Inter     uint64 `json:"inter"`
	MaxConn   uint64 `json:"maxconn"`
	MaxQueue  uint64 `json:"maxqueue"`
	Rise      uint64 `json:"rise"`
	SlowStart uint64 `json:"slowstart"`
	Weight    uint64 `json:"weight"`
}

type Server struct {
	Address        string         `json:"address"`
	Check          string         `json:"check"`
	GUID           string         `json:"guid"`
	Name           string         `json:"name"`
	Port           uint64         `json:"port"`
	ServerSettings ServerSettings `json:"serverSettings"`
}

type Backends struct {
	Algorithm             string         `json:"algorithm"`
	GUID                  string         `json:"guid"`
	Name                  string         `json:"name"`
	ServerDefaultSettings ServerSettings `json:"serverDefaultSettings"`
	Servers               []Server       `json:"servers"`
}

type Binding struct {
	Address string `json:"address"`
	GUID    string `json:"guid"`
	Name    string `json:"name"`
	Port    uint64 `json:"port"`
}

type Frontend struct {
	Backend  string    `json:"backend"`
	Bindings []Binding `json:"bindings"`
	GUID     string    `json:"guid"`
	Name     string    `json:"name"`
}

type PrimaryNode struct {
	BackendIP  string `json:"backendIp"`
	ComputeID  uint64 `json:"computeId"`
	FrontendIP string `json:"frontendIp"`
	GUID       string `json:"guid"`
	MgmtIP     string `json:"mgmtIp"`
	NetworkID  uint64 `json:"networkId"`
}

type SecondaryNode struct {
	BackendIP  string `json:"backendIp"`
	ComputeID  uint64 `json:"computeId"`
	FrontendIP string `json:"frontendIp"`
	GUID       string `json:"guid"`
	MgmtIP     string `json:"mgmtIp"`
	NetworkID  uint64 `json:"networkId"`
}

type LB struct {
	HAMode        bool          `json:"HAmode"`
	ACL           []ACL         `json:"acl"`
	Backends      []Backends    `json:"backends"`
	CreatedBy     string        `json:"createdBy"`
	CreatedTime   uint64        `json:"createdTime"`
	DeletedBy     string        `json:"deletedBy"`
	DeletedTime   uint64        `json:"deletedTime"`
	Desc          string        `json:"desc"`
	DpAPIUser     string        `json:"dpApiUser"`
	ExtNetID      uint64        `json:"extnetId"`
	Frontends     []Frontend    `json:"frontends"`
	GID           uint64        `json:"gid"`
	GUID          uint64        `json:"guid"`
	ID            uint64        `json:"id"`
	ImageID       uint64        `json:"imageId"`
	Milestones    uint64        `json:"milestones"`
	Name          string        `json:"name"`
	PrimaryNode   PrimaryNode   `json:"primaryNode"`
	RGID          uint64        `json:"rgId"`
	RGName        string        `json:"rgName"`
	SecondaryNode SecondaryNode `json:"secondaryNode"`
	Status        string        `json:"status"`
	TechStatus    string        `json:"techStatus"`
	UpdatedBy     string        `json:"updatedBy"`
	UpdatedTime   uint64        `json:"updatedTime"`
	VINSID        uint64        `json:"vinsId"`
}

type ListLB []LB

package lb

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
	ExtNetID      uint64      `json:"extnetId"`
	Frontends     []Frontend  `json:"frontends"`
	GID           uint64      `json:"gid"`
	GUID          uint64      `json:"guid"`
	ID            uint64      `json:"id"`
	ImageID       uint64      `json:"imageId"`
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
	VINSID        uint64      `json:"vinsId"`
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
	Rise      uint64 `json:"rise"`
	Fall      uint64 `json:"fall"`
	SlowStart uint64 `json:"slowstart"`
	MaxConn   uint64 `json:"maxconn"`
	MaxQueue  uint64 `json:"maxqueue"`
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

type Node struct {
	BackendIP  string `json:"backendIp"`
	ComputeID  uint64 `json:"computeId"`
	FrontendIP string `json:"frontendIp"`
	GUID       string `json:"guid"`
	MGMTIP     string `json:"mgmtIp"`
	NetworkID  uint64 `json:"networkId"`
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
	Port    uint64 `json:"port"`
}

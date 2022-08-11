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

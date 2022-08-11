package vins

type VinsRecord struct {
	AccountId   int    `json:"accountId"`
	AccountName string `json:"accountName"`
	CreatedBy   string `json:"createdBy"`
	CreatedTime int    `json:"createdTime"`
	DeletedBy   string `json:"deletedBy"`
	DeletedTime int    `json:"deletedTime"`
	ExternalIP  string `json:"externalIP"`
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Network     string `json:"network"`
	RGID        int    `json:"rgId"`
	RGName      string `json:"rgName"`
	Status      string `json:"status"`
	UpdatedBy   string `json:"updatedBy"`
	UpdatedTime int    `json:"updatedTime"`
	VXLanID     int    `json:"vxlanId"`
}

type VinsList []VinsRecord

type VinsAudits struct {
	Call         string  `json:"call"`
	ResponseTime float64 `json:"responsetime"`
	StatusCode   int     `json:"statuscode"`
	Timestamp    float64 `json:"timestamp"`
	User         string  `json:"user"`
}

type VinsAuditsList []VinsAudits

type VinsExtnet struct {
	DefaultGW  string `json:"default_gw"`
	ExtNetID   uint64 `json:"ext_net_id"`
	IP         string `json:"ip"`
	PrefixLen  uint   `json:"prefixlen"`
	Status     string `json:"status"`
	TechStatus string `json:"techStatus"`
}

type ExtnetList []VinsExtnet

type IP struct {
	ClientType string `json:"clientType"`
	DomainName string `json:"domainname"`
	HostName   string `json:"hostname"`
	IP         string `json:"ip"`
	MAC        string `json:"mac"`
	Type       string `json:"type"`
	VMId       uint64 `json:"vmId"`
}

type IPList []IP

type VNFDev struct {
	CKey            string           `json:"_ckey"`
	AccountId       uint64           `json:"accountId"`
	Capabilities    []string         `json:"capabilities"`
	Config          VNFConfig        `json:"config"`
	ConfigSaved     bool             `json:"configSaved"`
	CustomPreConfig bool             `json:"customPrecfg"`
	Description     string           `json:"desc"`
	GID             uint64           `json:"gid"`
	GUID            uint64           `json:"guid"`
	ID              uint64           `json:"id"`
	Interfaces      VNFInterfaceList `json:"interfaces"`
	LockStatus      string           `json:"lockStatus"`
	Milestones      uint64           `json:"milestones"`
	Name            string           `json:"name"`
	Status          string           `json:"status"`
	TechStatus      string           `json:"techStatus"`
	Type            string           `json:"type"`
	Vins            []uint64         `json:"vins"`
}

type VNFConfig struct {
	MGMT      VNFConfigMGMT      `json:"mgmt"`
	Resources VNFConfigResources `json:"resources"`
}

type VNFConfigMGMT struct {
	IPAddr   string `json:"ipaddr"`
	Password string `json:"password"`
	SSHKey   string `json:"sshkey"`
	User     string `json:"user"`
}

type VNFConfigResources struct {
	CPU     uint64 `json:"cpu"`
	RAM     uint64 `json:"ram"`
	StackId uint64 `json:"stackId"`
	UUID    string `json:"uuid"`
}

type VNFInterface struct {
	ConnId      uint64   `json:"connId"`
	ConnType    string   `json:"connType"`
	DefGW       string   `json:"defGw"`
	FlipGroupId uint64   `json:"flipgroupId"`
	GUID        string   `json:"guid"`
	IPAddress   string   `json:"ipAddress"`
	ListenSSH   bool     `json:"listenSsh"`
	MAC         string   `json:"mac"`
	Name        string   `json:"name"`
	NetId       uint64   `json:"netId"`
	NetMask     uint64   `json:"netMask"`
	NetType     string   `json:"netType"`
	PCISlot     uint64   `json:"pciSlot"`
	QOS         QOS      `json:"qos"`
	Target      string   `json:"target"`
	Type        string   `json:"type"`
	VNFS        []uint64 `json:"vnfs"`
}

type QOS struct {
	ERate   uint64 `json:"eRate"`
	GUID    string `json:"guid"`
	InBurst uint64 `json:"inBurst"`
	InRate  uint64 `json:"inRate"`
}

type VNFInterfaceList []VNFInterface

type VINSCompute struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type VINSComputeList []VINSCompute

type VNFS struct {
	DHCP DHCP `json:"DHCP"`
	GW   GW   `json:"GW"`
	NAT  NAT  `json:"NAT"`
}

type NAT struct {
	CKey        string  `json:"_ckey"`
	AccountId   uint64  `json:"accountId"`
	CreatedTime uint64  `json:"createdTime"`
	Devices     Devices `json:"devices"`
	GID         uint64  `json:"gid"`
	GUID        uint64  `json:"guid"`
	ID          uint64  `json:"id"`
	LockStatus  string  `json:"lockStatus"`
	Milestones  uint64  `json:"milestones"`
	OwnerId     uint64  `json:"ownerId"`
	OwnerType   string  `json:"ownerType"`
	PureVirtual bool    `json:"pureVirtual"`
	Status      string  `json:"status"`
	TechStatus  string  `json:"techStatus"`
	Type        string  `json:"type"`
}

type NATConfig struct {
	NetMask uint64        `json:"netmask"`
	Network string        `json:"network"`
	Rules   []interface{} `json:"rules"`
}

type GW struct {
	CKey        string   `json:"_ckey"`
	AccountId   uint64   `json:"accountId"`
	Config      GWConfig `json:"config"`
	CreatedTime uint64   `json:"createdTime"`
	Devices     Devices  `json:"devices"`
	GID         uint64   `json:"gid"`
	GUID        uint64   `json:"guid"`
	ID          uint64   `json:"id"`
	LockStatus  string   `json:"lockStatus"`
	Milestones  uint64   `json:"milestones"`
	OwnerId     uint64   `json:"ownerId"`
	OwnerType   string   `json:"ownerType"`
	PureVirtual bool     `json:"pureVirtual"`
	Status      string   `json:"status"`
	TechStatus  string   `json:"techStatus"`
	Type        string   `json:"type"`
}

type GWConfig struct {
	DefaultGW  string `json:"default_gw"`
	ExtNetId   uint64 `json:"ext_net_id"`
	ExtNetIp   string `json:"ext_net_ip"`
	ExtNetMask uint64 `json:"ext_netmask"`
	QOS        QOS    `json:"qos"`
}

type Devices struct {
	Primary DevicePrimary `json:"primary"`
}

type DevicePrimary struct {
	DevId   uint64 `json:"devId"`
	IFace01 string `json:"iface01"`
	IFace02 string `json:"iface02"`
}

type DHCP struct {
	CKey        string     `json:"_ckey"`
	AccountId   uint64     `json:"accountId"`
	Config      DHCPConfig `json:"config"`
	CreatedTime uint64     `json:"createdTime"`
	Devices     Devices    `json:"devices"`
	GID         uint64     `json:"gid"`
	GUID        uint64     `json:"guid"`
	ID          uint64     `json:"id"`
	LockStatus  string     `json:"lockStatus"`
	Milestones  uint64     `json:"milestones"`
	OwnerId     uint64     `json:"ownerId"`
	OwnerType   string     `json:"ownerType"`
	PureVirtual bool       `json:"pureVirtual"`
	Status      string     `json:"status"`
	TechStatus  string     `json:"techStatus"`
	Type        string     `json:"type"`
}

type DHCPConfig struct {
	DefaultGW    string          `json:"default_gw"`
	DNS          []string        `json:"dns"`
	IPEnd        string          `json:"ip_end"`
	IPStart      string          `json:"ip_start"`
	Lease        uint64          `json:"lease"`
	Netmask      uint64          `json:"netmask"`
	Network      string          `json:"network"`
	Reservations ReservationList `json:"reservations"`
}

type VinsDetailed struct {
	VNFDev            VNFDev          `json:"VNFDev"`
	CKey              string          `json:"_ckey"`
	AccountId         uint64          `json:"accountId"`
	AccountName       string          `json:"accountName"`
	Computes          VINSComputeList `json:"computes"`
	DefaultGW         string          `json:"defaultGW"`
	DefaultQOS        QOS             `json:"defaultQos"`
	Description       string          `json:"desc"`
	GID               uint64          `json:"gid"`
	GUID              uint64          `json:"guid"`
	ID                uint64          `json:"id"`
	LockStatus        string          `json:"lockStatus"`
	ManagerId         uint64          `json:"managerId"`
	ManagerType       string          `json:"managerType"`
	Milestones        uint64          `json:"milestones"`
	Name              string          `json:"name"`
	NetMask           uint64          `json:"netMask"`
	Network           string          `json:"network"`
	PreReservaionsNum uint64          `json:"preReservationsNum"`
	Redundant         bool            `json:"redundant"`
	RGID              uint64          `json:"rgId"`
	RGName            string          `json:"rgName"`
	SecVNFDevId       uint64          `json:"secVnfDevId"`
	Status            string          `json:"status"`
	UserManaged       bool            `json:"userManaged"`
	VNFS              VNFS            `json:"vnfs"`
	VXLanId           uint64          `json:"vxlanId"`
}

type Reservation struct {
	ClientType  string `json:"clientType"`
	Description string `json:"desc"`
	DomainName  string `json:"domainname"`
	HostName    string `json:"hostname"`
	IP          string `json:"ip"`
	MAC         string `json:"mac"`
	Type        string `json:"type"`
	VMID        int    `json:"vmId"`
}

type ReservationList []Reservation

package vins

// Main information about audit
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

// Main information about external network
type ItemExtNet struct {
	// Default GW
	DefaultGW string `json:"default_gw"`

	// External network ID
	ExtNetID uint64 `json:"ext_net_id"`

	// IP
	IP string `json:"ip"`

	// Prefix len
	PrefixLen uint64 `json:"prefixlen"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`
}

// List of external networks
type ListExtNets []ItemExtNet

// MGMT
type MGMT struct {
	// IP address
	IPAddress string `json:"ipaddr"`

	// Password
	Password string `json:"password"`

	// SSH key
	SSHKey string `json:"sshkey"`

	// User
	User string `json:"user"`
}

// Main information about resource
type Resources struct {
	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// Stack ID
	StackID uint64 `json:"stackId"`

	// UUID
	UUID string `json:"uuid"`
}

// VNF config
type Config struct {
	// MGMT
	MGMT MGMT `json:"mgmt"`

	// Resources
	Resources Resources `json:"resources"`
}

// Main information about QOS
type QOS struct {
	// ERate
	ERate uint64 `json:"eRate"`

	// GUID
	GUID string `json:"guid"`

	// InBurst
	InBurst uint64 `json:"inBurst"`

	// InRate
	InRate uint64 `json:"inRate"`
}

// Main information about interface
type ItemInterface struct {
	// Connection ID
	ConnID uint64 `json:"connId"`

	// Connection type
	ConnType string `json:"connType"`

	// Default GW
	DefGW string `json:"defGw"`

	// FLIPGroup ID
	FLIPGroupID uint64 `json:"flipgroupId"`

	// GUID
	GUID string `json:"guid"`

	// IP address
	IPAddress string `json:"ipAddress"`

	// Listen SSH
	ListenSSH bool `json:"listenSsh"`

	// MAC
	MAC string `json:"mac"`

	// Name
	Name string `json:"name"`

	// Network type
	NetID uint64 `json:"netId"`

	// Network mask
	NetMask uint64 `json:"netMask"`

	// Network type
	NetType string `json:"netType"`

	// PCI slot
	PCISlot uint64 `json:"pciSlot"`

	// QOS
	QOS QOS `json:"qos"`

	// Target
	Target string `json:"target"`

	// Type
	Type string `json:"type"`

	// List of VNF IDs
	VNFs []uint64 `json:"vnfs"`
}

// List of interfaces
type ListInterfaces []ItemInterface

// Main information about VNF device
type VNFDev struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Capabilities
	Capabilities []string `json:"capabilities"`

	// Config
	Config Config `json:"config"`

	// Config saved
	ConfigSaved bool `json:"configSaved"`

	// CustomPreConfig
	CustomPreConfig bool `json:"customPrecfg"`

	// Description
	Description string `json:"desc"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// List of interfaces
	Interfaces ListInterfaces `json:"interfaces"`

	// Lock status
	LockStatus string `json:"lockStatus"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// TechStatus
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`

	//List of VINS IDs
	VINS []uint64 `json:"vins"`
}

// Main information about reservation
type ItemReservation struct {
	// Client type
	ClientType string `json:"clientType"`

	// Description
	Description string `json:"desc"`

	// Domain name
	DomainName string `json:"domainname"`

	// Hostname
	Hostname string `json:"hostname"`

	// IP
	IP string `json:"ip"`

	// MAC
	MAC string `json:"mac"`

	// Type
	Type string `json:"type"`

	// Virtual machine ID
	VMID uint64 `json:"vmId"`
}

// List of reservations
type ListReservations []ItemReservation

// VNFs sonfig
type VNFsConfig struct {
	// Default GW
	DefaultGW string `json:"default_gw"`

	// List of DNS
	DNS []string `json:"dns"`

	// IP end
	IPEnd string `json:"ip_end"`

	// IP start
	IPStart string `json:"ip_start"`

	// Lease
	Lease uint64 `json:"lease"`

	// Network mask
	NetMask uint64 `json:"netmask"`

	// Network
	Network string `json:"network"`

	// List of reservations
	Reservations ListReservations `json:"reservations"`
}

// Primary
type Primary struct {
	// Device ID
	DevID uint64 `json:"devId"`

	// IFace01
	IFace01 string `json:"iface01"`

	// IFace02
	IFace02 string `json:"iface02"`
}

// Devices
type Devices struct {
	// Primary
	Primary Primary `json:"primary"`
}

// Main information about DHCP
type RecordDHCP struct {
	// Config
	Config VNFsConfig `json:"config"`

	// DHCP details
	InfoVNF
}

// GW config
type GWConfig struct {
	// Default GW
	DefaultGW string `json:"default_gw"`

	// External network ID
	ExtNetID uint64 `json:"ext_net_id"`

	// External network IP
	ExtNetIP string `json:"ext_net_ip"`

	// External network mask
	ExtNetMask uint64 `json:"ext_netmask"`

	// QOS
	QOS QOS `json:"qos"`
}

// Main information about GW
type RecordGW struct {
	// Config
	Config GWConfig `json:"config"`

	// GW details
	InfoVNF
}

// NAT config
type NATConfig struct {
	// Network mask
	NetMask uint64 `json:"netmask"`

	// Network
	Network string `json:"network"`

	// Rules
	Rules []interface{} `json:"rules"`
}

// Main information about NAT
type RecordNAT struct {
	// Config
	Config NATConfig `json:"config"`

	// NAT details
	InfoVNF
}

// NAT/GW/NAT details
type InfoVNF struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Config
	Config NATConfig `json:"config"`

	// CreatedTime
	CreatedTime uint64 `json:"createdTime"`

	// Devices
	Devices Devices `json:"devices"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Lock status
	LockStatus string `json:"lockStatus"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Owner ID
	OwnerID uint64 `json:"ownerId"`

	// Owner type
	OwnerType string `json:"ownerType"`

	// Pure virtual
	PureVirtual bool `json:"pureVirtual"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`
}

// main information about VNF
type RecordVNFs struct {
	// DHCP
	DHCP RecordDHCP `json:"DHCP"`

	// GW
	GW RecordGW `json:"GW"`

	// NAT
	NAT RecordNAT `json:"NAT"`
}

// Detailed information about VINS
type RecordVINS struct {
	// VNF device
	VNFDev VNFDev `json:"VNFDev"`

	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Default GW
	DefaultGW string `json:"defaultGW"`

	// Default QOS
	DefaultQOS QOS `json:"defaultQos"`

	// Description
	Description string `json:"desc"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Lock status
	LockStatus string `json:"lockStatus"`

	// Manager ID
	ManagerID uint64 `json:"managerId"`

	// Manager type
	ManagerType string `json:"managerType"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Network mask
	NetMask uint64 `json:"netMask"`

	// Network
	Network string `json:"network"`

	// PreReservationsNum
	PreReservationsNum uint64 `json:"preReservationsNum"`

	// Redundant
	Redundant bool `json:"redundant"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// SecVNFDevID
	SecVNFDevID uint64 `json:"secVnfDevId"`

	// Status
	Status string `json:"status"`

	// User managed
	UserManaged bool `json:"userManaged"`

	// VNFs information
	VNFs RecordVNFs `json:"vnfs"`

	// VXLAN ID
	VXLANID uint64 `json:"vxlanId"`
}

// Main information about IP
type ItemIP struct {
	// IP
	IP string `json:"ip"`

	// MAC
	MAC string `json:"mac"`

	// Type
	Type string `json:"type"`

	// Virtual machine ID
	VMID uint64 `json:"vmId"`

	// Client type
	ClientType string `json:"clientType"`

	// Domain name
	DomainName string `json:"domainname"`

	// Hostname
	Hostname string `json:"hostname"`
}

// List of information about IPs
type ListIPs []ItemIP

// Main information about NAT rule
type ItemNATRule struct {
	// ID
	ID uint64 `json:"id"`

	// Local IP
	LocalIP string `json:"localIp"`

	// Local port
	LocalPort uint64 `json:"localPort"`

	// Protocol
	Protocol string `json:"protocol"`

	// Public port end
	PublicPortEnd uint64 `json:"publicPortEnd"`

	// Public port start
	PublicPortStart uint64 `json:"publicPortStart"`

	// Virtual machine ID
	VMID uint64 `json:"vmId"`

	// Virtual machine name
	VMName string `json:"vmName"`
}

// List NAT rules
type ListNATRules []ItemNATRule

// Shorted information about VNF
type ItemVNFs struct {
	// DHCP
	DHCP uint64 `json:"dhcp"`

	// DNS
	DNS uint64 `json:"dns"`

	// FW
	FW uint64 `json:"fw"`

	// GW
	GW uint64 `json:"gw"`

	// NAT
	NAT uint64 `json:"nat"`

	// VPN
	VPN uint64 `json:"vpn"`
}

// Main information about VINS
type ItemVINS struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Default GW
	DefaultGW string `json:"defaultGW"`

	// Default QOS
	DefaultQOS QOS `json:"defaultQos"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// External IP
	ExternalIP string `json:"externalIP"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Lock status
	LockStatus string `json:"lockStatus"`

	// Manager ID
	ManagerID uint64 `json:"managerId"`

	// Manager type
	ManagerType string `json:"managerType"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Network mask
	NetMask uint64 `json:"netMask"`

	// Network
	Network string `json:"network"`

	// PreReservationsNum
	PreReservationsNum uint64 `json:"preReservationsNum"`

	// PriVNFDevID
	PriVNFDevID uint64 `json:"priVnfDevId"`

	// Redundant
	Redundant bool `json:"redundant"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// SecVNFDevID
	SecVNFDevID uint64 `json:"secVnfDevId"`

	// Status
	Status string `json:"status"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// User managed
	UserManaged bool `json:"userManaged"`

	// VNFs
	VNFs ItemVNFs `json:"vnfs"`

	// VXLAN ID
	VXLANID uint64 `json:"vxlanId"`
}

// List of VINS
type ListVINS []ItemVINS

package vins

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

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// External IP
	ExternalIP string `json:"externalIP"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Network
	Network string `json:"network"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// Status
	Status string `json:"status"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// VXLAN ID
	VXLANID uint64 `json:"vxlanId"`
}

// List of VINSes
type ListVINS []ItemVINS

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

// Main information about IP
type ItemIP struct {
	// Client type
	ClientType string `json:"clientType"`

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

// List of IPs
type ListIPs []ItemIP

// Main information about VNF device
type RecordVNFDev struct {
	// CKey
	CKey string `json:"_ckey"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Capabilities
	Capabilities []string `json:"capabilities"`

	// Config
	Config RecordVNFConfig `json:"config"`

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
	Interfaces ListVNFInterfaces `json:"interfaces"`

	// Lock status
	LockStatus string `json:"lockStatus"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`

	// List of VINS IDs
	VINS []uint64 `json:"vins"`
}

// VNF config
type RecordVNFConfig struct {
	// MGMT
	MGMT RecordMGMT `json:"mgmt"`

	// Resources
	Resources RecordResources `json:"resources"`
}

// Main information about MGMT
type RecordMGMT struct {
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
type RecordResources struct {
	// Number of CPU
	CPU uint64 `json:"cpu"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// Stack ID
	StackID uint64 `json:"stackId"`

	// UUID
	UUID string `json:"uuid"`
}

// Main information about VNF interface
type ItemVNFInterface struct {
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

// List of VNF interfaces
type ListVNFInterfaces []ItemVNFInterface

// Main information about VINS compute
type ItemVINSCompute struct {
	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`
}

// List of VINS computes
type ListVINSComputes []ItemVINSCompute

// Detailed information about VNF
type RecordVNFs struct {
	// DHCP
	DHCP RecordDHCP `json:"DHCP"`

	// GW
	GW RecordGW `json:"GW"`

	// NAT
	NAT RecordNAT `json:"NAT"`
}

// Main information about NAT
type RecordNAT struct {
	// CKey
	CKey string `json:"_ckey"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Config
	Config NATConfig `json:"config"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Detailed information about devices
	Devices RecordDevices `json:"devices"`

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

// NAT configuration
type NATConfig struct {
	// Network mask
	NetMask uint64 `json:"netmask"`

	// Network
	Network string `json:"network"`

	// List NAT rules
	Rules ListNATRules `json:"rules"`
}

// Main information about GW
type RecordGW struct {
	// CKey
	CKey string `json:"_ckey"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Config
	Config RecordGWConfig `json:"config"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Detailed information about devices
	Devices RecordDevices `json:"devices"`

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

// GW configuration
type RecordGWConfig struct {
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

// Information about devices
type RecordDevices struct {
	// Main information about primary device
	Primary RecordPrimary `json:"primary"`
}

// Main information about primary device
type RecordPrimary struct {
	// Device ID
	DevID uint64 `json:"devId"`

	// IFace01
	IFace01 string `json:"iface01"`

	// IFace02
	IFace02 string `json:"iface02"`
}

// Main information about DHCP
type RecordDHCP struct {
	// CKey
	CKey string `json:"_ckey"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Config
	Config RecordDHCPConfig `json:"config"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Detailed information about devices
	Devices RecordDevices `json:"devices"`

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

// DHCP configuration
type RecordDHCPConfig struct {
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

// Detailed information about VINS
type RecordVINS struct {
	// Main information about VNF device
	VNFDev RecordVNFDev `json:"VNFDev"`

	// CKey
	CKey string `json:"_ckey"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// List of VINS computes
	Computes ListVINSComputes `json:"computes"`

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

	// Pre reservaions number
	PreReservaionsNum uint64 `json:"preReservationsNum"`

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

	// User managed
	UserManaged bool `json:"userManaged"`

	// Main information about VNFs
	VNFs RecordVNFs `json:"vnfs"`

	//  VXLAN ID
	VXLANID uint64 `json:"vxlanId"`
}

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

// List of NAT rules
type ListNATRules []ItemNATRule

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

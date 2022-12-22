package extnet

// Main information about external network
type ItemExtNet struct {
	// ID
	ID uint64 `json:"id"`

	// IPCIDR
	IPCIDR string `json:"ipcidr"`

	// Name
	Name string `json:"name"`

	// Status
	Status string `json:"status"`
}

// Extend information about external network
type ItemExtNetExtend struct {
	// Main information about external network
	ItemExtNet

	// IP address
	IPAddr string `json:"ipaddr"`
}

// List of information about external network
type ListExtNets []ItemExtNet

// List of extend information about external network
type ListExtNetExtends []ItemExtNetExtend

// Main information about compute with external network
type ItemExtNetCompute struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// List of extend information about external network
	ExtNets ListExtNetExtends `json:"extnets"`

	// ID
	ID uint64 `json:"id"`

	// Name
	Name string `json:"name"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`
}

// List of information about computes with external network
type ListExtNetComputes []ItemExtNetCompute

// QOS
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

// Main information about reservations
type ItemReservation struct {
	// ClientType
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

// List of information about reservations
type ListReservations []ItemReservation

// VNFs
type VNFs struct {
	DHCP uint64 `json:"dhcp"`
}

// Detailed information about external network
type RecordExtNet struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// CheckIPs
	CheckIPs []string `json:"checkIPs"`

	// CheckIps
	CheckIps []string `json:"checkIps"`

	// Default
	Default bool `json:"default"`

	// Default QOS
	DefaultQOS QOS `json:"defaultQos"`

	// Description
	Description string `json:"desc"`

	// list of DNS
	DNS []string `json:"dns"`

	// Excluded
	Excluded []string `json:"excluded"`

	// Free IPs
	FreeIPs uint64 `json:"free_ips"`

	// Gateway
	Gateway string `json:"gateway"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// IPCIDR
	IPCIDR string `json:"ipcidr"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Network
	Network string `json:"network"`

	// Network ID
	NetworkID uint64 `json:"networkId"`

	// OVS Bridge
	OVSBridge string `json:"ovsBridge"`

	// PreReservation IP num
	PreReservationsNum uint64 `json:"preReservationsNum"`

	// Prefix
	Prefix uint64 `json:"prefix"`

	// PriVNFDevID
	PriVNFDevID uint64 `json:"priVnfDevId"`

	// List reservations
	Reservations ListReservations `json:"reservations"`

	// Shared with
	SharedWith []uint64 `json:"sharedWith"`

	// Status
	Status string `json:"status"`

	// VLAN ID
	VLANID uint64 `json:"vlanId"`

	// VNFs
	VNFs VNFs `json:"vnfs"`
}

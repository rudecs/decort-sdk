package extnet

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
	// Client type
	ClientType string `json:"clientType"`

	// Description
	Description string `json:"desc"`

	// IP
	IP string `json:"ip"`

	// MAC
	MAC string `json:"mac"`

	// Type
	Type string `json:"type"`

	// Virtual machine ID
	VMID uint64 `json:"vmId"`

	// Domain name
	DomainName string `json:"domainname"`

	// Hostname
	Hostname string `json:"hostname"`
}

// List reservations
type ListReservations []ItemReservation

// VNFs
type VNFs struct {
	DHCP int `json:"dhcp"`
}

// Main information about external network
type ItemExtNet struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// CheckIPs
	CheckIPs []string `json:"checkIps"`

	// Default
	Default bool `json:"default"`

	// Default QOS
	DefaultQOS QOS `json:"defaultQos"`

	// Description
	Description string `json:"desc"`

	// Free IPs number
	FreeIPs uint64 `json:"freeIps"`

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

	// Network ID
	NetworkID uint64 `json:"networkId"`

	// OVSBridge
	OVSBridge string `json:"ovsBridge"`

	// PreReservationsNum
	PreReservationsNum uint64 `json:"preReservationsNum"`

	// PriVNFDevID
	PriVNFDevID uint64 `json:"priVnfDevId"`

	// List of shared with
	SharedWith []interface{} `json:"sharedWith"`

	// Status
	Status string `json:"status"`

	// VLAN ID
	VLANID uint64 `json:"vlanId"`

	// VNFs
	VNFs VNFs `json:"vnfs"`
}

// List external networks
type ListExtNet []ItemExtNet

// Detailed information about external network
type RecordExtNet struct {
	// Main information about external network
	ItemExtNet

	// CheckIPs
	CheckIPs []string `json:"checkIPs"`

	// List DNS
	DNS []string `json:"dns"`

	// List excludes
	Excluded []string `json:"excluded"`

	// Gateway
	Gateway string `json:"gateway"`

	// Network
	Network string `json:"network"`

	// Prefix
	Prefix uint64 `json:"prefix"`

	// List reservations
	Reservations ListReservations `json:"reservations"`
}

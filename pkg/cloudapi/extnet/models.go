package extnet

type ExtNetRecord struct {
	ID     uint64 `json:"id"`
	IPCidr string `json:"ipcidr"`
	Name   string `json:"name"`
}
type ExtNetExtend struct {
	ExtNetRecord
	IPAddr string `json:"ipaddr"`
}

type ExtNetList []ExtNetRecord
type ExtNetExtendList []ExtNetExtend

type ExtNetComputes struct {
	AccountID   uint64           `json:"accountId"`
	AccountName string           `json:"accountName"`
	ExtNets     ExtNetExtendList `json:"extnets"`
	ID          uint64           `json:"id"`
	Name        string           `json:"name"`
	RGID        uint64           `json:"rgId"`
	RGName      string           `json:"rgName"`
}

type ExtNetComputesList []ExtNetComputes

type ExtNetQos struct {
	ERate   uint64 `json:"eRate"`
	GUID    string `json:"guid"`
	InBurst uint64 `json:"inBurst"`
	InRate  uint64 `json:"inRate"`
}

type ExtNetReservation struct {
	ClientType  string `json:"clientType"`
	Description string `json:"desc"`
	DomainName  string `json:"domainname"`
	HostName    string `json:"hostname"`
	IP          string `json:"ip"`
	MAC         string `json:"mac"`
	Type        string `json:"type"`
	VMID        uint64 `json:"vmId"`
}

type ExtNetReservations []ExtNetReservation

type ExtNetVNFS struct {
	DHCP uint64 `json:"dhcp"`
}

type ExtNetDetailed struct {
	CKey               string             `json:"_ckey"`
	Meta               []interface{}      `json:"_meta"`
	CheckIPs           []string           `json:"checkIPs"`
	CheckIps           []string           `json:"checkIps"`
	Default            bool               `json:"default"`
	DefaultQos         ExtNetQos          `json:"defaultQos"`
	Description        string             `json:"desc"`
	Dns                []string           `json:"dns"`
	Excluded           []string           `json:"excluded"`
	FreeIps            uint64             `json:"free_ips"`
	Gateway            string             `json:"gateway"`
	GID                uint64             `json:"gid"`
	GUID               uint64             `json:"guid"`
	ID                 uint64             `json:"id"`
	IPCidr             string             `json:"ipcidr"`
	Milestones         uint64             `json:"milestones"`
	Name               string             `json:"name"`
	Network            string             `json:"network"`
	NetworkID          uint64             `json:"networkId"`
	PreReservationsNum uint64             `json:"preReservationsNum"`
	Prefix             uint64             `json:"prefix"`
	PriVNFDevID        uint64             `json:"priVnfDevId"`
	Reservations       ExtNetReservations `json:"reservations"`
	SharedWith         []uint64           `json:"sharedWith"`
	Status             string             `json:"status"`
	VlanID             uint64             `json:"vlanId"`
	VNFS               ExtNetVNFS         `json:"vnfs"`
}

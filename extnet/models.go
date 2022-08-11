package extnet

type ExtnetRecord struct {
	ID     int    `json:"id"`
	IPCidr string `json:"ipcidr"`
	Name   string `json:"name"`
}
type ExtnetExtend struct {
	ExtnetRecord
	IPAddr string `json:"ipaddr"`
}

type ExtnetList []ExtnetRecord
type ExtnetExtendList []ExtnetExtend

type ExtnetComputes struct {
	AccountId   int              `json:"accountId"`
	AccountName string           `json:"accountName"`
	Extnets     ExtnetExtendList `json:"extnets"`
	ID          int              `json:"id"`
	Name        string           `json:"name"`
	RGID        int              `json:"rgId"`
	RGName      string           `json:"rgName"`
}

type ExtnetComputesList []ExtnetComputes

type ExtnetQos struct {
	ERate   int    `json:"eRate"`
	GUID    string `json:"guid"`
	InBurst int    `json:"inBurst"`
	InRate  int    `json:"inRate"`
}

type ExtnetReservation struct {
	ClientType  string `json:"clientType"`
	Description string `json:"desc"`
	DomainName  string `json:"domainname"`
	HostName    string `json:"hostname"`
	IP          string `json:"ip"`
	MAC         string `json:"mac"`
	Type        string `json:"type"`
	VMID        int    `json:"vmId"`
}

type ExtnetReservations []ExtnetReservation

type ExtnetVNFS struct {
	DHCP int `json:"dhcp"`
}

type ExtnetDetailed struct {
	CKey               string             `json:"_ckey"`
	Meta               []interface{}      `json:"_meta"`
	CheckIPs           []string           `json:"checkIPs"`
	CheckIps           []string           `json:"checkIps"`
	Default            bool               `json:"default"`
	DefaultQos         ExtnetQos          `json:"defaultQos"`
	Description        string             `json:"desc"`
	Dns                []string           `json:"dns"`
	Excluded           []string           `json:"excluded"`
	FreeIps            int                `json:"free_ips"`
	Gateway            string             `json:"gateway"`
	GID                int                `json:"gid"`
	GUID               int                `json:"guid"`
	ID                 int                `json:"id"`
	IPCidr             string             `json:"ipcidr"`
	Milestones         int                `json:"milestones"`
	Name               string             `json:"name"`
	Network            string             `json:"network"`
	NetworkId          int                `json:"networkId"`
	PreReservationsNum int                `json:"preReservationsNum"`
	Prefix             int                `json:"prefix"`
	PriVnfDevId        int                `json:"priVnfDevId"`
	Reservations       ExtnetReservations `json:"reservations"`
	SharedWith         []int              `json:"sharedWith"`
	Status             string             `json:"status"`
	VlanID             int                `json:"vlanId"`
	VNFS               ExtnetVNFS         `json:"vnfs"`
}

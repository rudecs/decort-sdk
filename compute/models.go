package compute

//ACL for compute
type UserList struct {
	AccountACL ACLList `json:"accountACL"`
	ComputeACL ACLList `json:"computeACL"`
	RGACL      ACLList `json:"rgACL"`
}

type ACL struct {
	Explicit    bool   `json:"explicit"`
	GUID        string `json:"guid"`
	Rigth       string `json:"right"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	UserGroupId string `json:"userGroupId"`
}

type ACLList []ACL

type SnapshotUsage struct {
	Count     uint64  `json:"count,omitempty"`
	Stored    float64 `json:"stored"`
	Label     string  `json:"label,omitempty"`
	Timestamp uint64  `json:"timestamp,omitempty"`
}

type SnapshotUsageList []SnapshotUsage

type Snapshot struct {
	Disks     []uint64 `json:"disks"`
	GUID      string   `json:"guid"`
	Label     string   `json:"label"`
	Timestamp uint64   `json:"timestamp"`
}

type SnapshotList []Snapshot

type PFW struct {
	ID              uint64 `json:"id"`
	LocalIP         string `json:"localIp"`
	LocalPort       uint64 `json:"localPort"`
	Protocol        string `json:"protocol"`
	PublicPortEnd   uint64 `json:"publicPortEnd"`
	PublicPortStart uint64 `json:"publicPortStart"`
	VMID            uint64 `json:"vmId"`
}

type PFWList []PFW

type AffinityRelations struct {
	OtherNode             []interface{} `json:"otherNode"`
	OtherNodeIndirect     []interface{} `json:"otherNodeIndirect"`
	OtherNodeIndirectSoft []interface{} `json:"otherNodeIndirectSoft"`
	OtherNodeSoft         []interface{} `json:"otherNodeSoft"`
	SameNode              []interface{} `json:"sameNode"`
	SameNodeSoft          []interface{} `json:"sameNodeSoft"`
}

type NetAttach struct {
	ConnID      uint64   `json:"connId"`
	ConnType    string   `json:"connType"`
	DefGW       string   `json:"defGw"`
	FlipGroupID uint64   `json:"flipgroupId"`
	GUID        string   `json:"guid"`
	IPAddress   string   `json:"ipAddress"`
	ListenSSH   bool     `json:"listenSsh"`
	MAC         string   `json:"mac"`
	Name        string   `json:"name"`
	NetID       uint64   `json:"netId"`
	NetMask     uint64   `json:"netMask"`
	NetType     string   `json:"netType"`
	PCISlot     uint64   `json:"pciSlot"`
	QOS         QOS      `json:"qos"`
	Target      string   `json:"target"`
	Type        string   `json:"type"`
	VNFS        []uint64 `json:"vnfs"`
}

type Audit struct {
	Call         string  `json:"call"`
	ResponseTime float64 `json:"responsetime"`
	StatusCode   uint64  `json:"statuscode"`
	Timestamp    float64 `json:"timestamp"`
	User         string  `json:"user"`
}

type AuditList []Audit

type AuditShort struct {
	Epoch   float64 `json:"epoch"`
	Message string  `json:"message"`
}

type AuditShortList []AuditShort

type Rule struct {
	GUID     string `json:"guid"`
	Key      string `json:"key"`
	Mode     string `json:"mode"`
	Policy   string `json:"policy"`
	Topology string `json:"topology"`
	Value    string `json:"value"`
}

type RuleList []Rule

type ComputeRecord struct {
	ACL                    UserList               `json:"ACL"`
	AccountID              uint64                 `json:"accountId"`
	AccountName            string                 `json:"accountName"`
	AffinityLabel          string                 `json:"affinityLabel"`
	AffinityRules          RuleList               `json:"affinityRules"`
	AffinityWeight         uint64                 `json:"affinityWeight"`
	AntiAffinityRules      RuleList               `json:"antiAffinityRules"`
	Architecture           string                 `json:"arch"`
	BootOrder              []string               `json:"bootOrder"`
	BootDiskSize           uint64                 `json:"bootdiskSize"`
	CloneReference         uint64                 `json:"cloneReference"`
	Clones                 []uint64               `json:"clones"`
	ComputeCIID            uint64                 `json:"computeciId"`
	CPU                    uint64                 `json:"cpus"`
	CreatedBy              string                 `json:"createdBy"`
	CreatedTime            uint64                 `json:"createdTime"`
	CustomFields           map[string]interface{} `json:"customFields"`
	DeletedBy              string                 `json:"deletedBy"`
	DeletedTime            uint64                 `json:"deletedTime"`
	Description            string                 `json:"desc"`
	Devices                interface{}            `json:"devices"`
	Disks                  ComputeDiskList        `json:"disks"`
	Driver                 string                 `json:"driver"`
	GID                    uint64                 `json:"gid"`
	GUID                   uint64                 `json:"guid"`
	ID                     uint64                 `json:"id"`
	ImageID                uint64                 `json:"imageId"`
	ImageName              string                 `json:"imageName"`
	Intefaces              IntefaceList           `json:"interfaces"`
	LockStatus             string                 `json:"lockStatus"`
	ManagerID              uint64                 `json:"managerId"`
	ManagerType            string                 `json:"managerType"`
	MigrationJob           uint64                 `json:"migrationjob"`
	Milestones             uint64                 `json:"milestones"`
	Name                   string                 `json:"name"`
	NatableVinsID          uint64                 `json:"natableVinsId"`
	NatableVinsIP          string                 `json:"natableVinsIp"`
	NatableVinsName        string                 `json:"natableVinsName"`
	NatableVinsNetwork     string                 `json:"natableVinsNetwork"`
	NatableVinsNetworkName string                 `json:"natableVinsNetworkName"`
	OSUsers                OSUserList             `json:"osUsers"`
	Pinned                 bool                   `json:"pinned"`
	RAM                    uint64                 `json:"ram"`
	ReferenceID            string                 `json:"referenceId"`
	Registered             bool                   `json:"registered"`
	ResName                string                 `json:"resName"`
	RGID                   uint64                 `json:"rgId"`
	RGName                 string                 `json:"rgName"`
	SnapSets               SnapSetList            `json:"snapSets"`
	StatelessSepID         uint64                 `json:"statelessSepId"`
	StatelessSepType       string                 `json:"statelessSepType"`
	Status                 string                 `json:"status"`
	Tags                   map[string]string      `json:"tags"`
	TechStatus             string                 `json:"techStatus"`
	UpdatedBy              string                 `json:"updatedBy"`
	UpdatedTime            uint64                 `json:"updatedTime"`
	UserManaged            bool                   `json:"userManaged"`
	Userdata               interface{}            `json:"userdata"`
	VGPUs                  []uint64               `json:"vgpus"`
	VirtualImageID         uint64                 `json:"virtualImageId"`
	VirtualImageName       string                 `json:"virtualImageName"`
}

type ComputeList []ComputeRecord

type OSUser struct {
	GUID     string `json:"guid"`
	Login    string `json:"login"`
	Password string `json:"password"`
	PubKey   string `json:"pubkey"`
}

type OSUserList []OSUser

type SnapSet struct {
	Disks     []uint64 `json:"disks"`
	GUID      string   `json:"guid"`
	Label     string   `json:"label"`
	Timestamp uint64   `json:"timestamp"`
}

type SnapSetList []SnapSet

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

type IntefaceList []VNFInterface

type ComputeDiskList []ComputeDisk

type ComputeDisk struct {
	Ckey                string                 `json:"_ckey"`
	Acl                 map[string]interface{} `json:"acl"`
	AccountID           int                    `json:"accountId"`
	Bootpartition       uint64                 `json:"bootPartition"`
	CreatedTime         uint64                 `json:"createdTime"`
	DeletedTime         uint64                 `json:"deletedTime"`
	Description         string                 `json:"desc"`
	DestructionTime     uint64                 `json:"destructionTime"`
	DiskPath            string                 `json:"diskPath"`
	GID                 uint64                 `json:"gid"`
	GUID                uint64                 `json:"guid"`
	ID                  uint64                 `json:"id"`
	ImageID             uint64                 `json:"imageId"`
	Images              []uint64               `json:"images"`
	IOTune              IOTune                 `json:"iotune"`
	IQN                 string                 `json:"iqn"`
	Login               string                 `json:"login"`
	Milestones          uint64                 `json:"milestones"`
	Name                string                 `json:"name"`
	Order               uint64                 `json:"order"`
	Params              string                 `json:"params"`
	ParentID            uint64                 `json:"parentId"`
	Passwd              string                 `json:"passwd"`
	PciSlot             uint64                 `json:"pciSlot"`
	Pool                string                 `json:"pool"`
	PurgeTime           uint64                 `json:"purgeTime"`
	RealityDeviceNumber uint64                 `json:"realityDeviceNumber"`
	ResID               string                 `json:"resId"`
	Role                string                 `json:"role"`
	SepID               int                    `json:"sepId"` // NOTE: absent from compute/get output
	SizeMax             int                    `json:"sizeMax"`
	SizeUsed            int                    `json:"sizeUsed"` // sum over all snapshots of this disk to report total consumed space
	Snapshots           SnapshotExtendList     `json:"snapshots"`
	Status              string                 `json:"status"`
	TechStatus          string                 `json:"techStatus"`
	Type                string                 `json:"type"`
	VMID                int                    `json:"vmid"`
}

type SnapshotExtend struct {
	Guid        string `json:"guid"`
	Label       string `json:"label"`
	ResId       string `json:"resId"`
	SnapSetGuid string `json:"snapSetGuid"`
	SnapSetTime uint64 `json:"snapSetTime"`
	TimeStamp   uint64 `json:"timestamp"`
}

type SnapshotExtendList []SnapshotExtend

type IOTune struct {
	ReadBytesSec     int `json:"read_bytes_sec"`
	ReadBytesSecMax  int `json:"read_bytes_sec_max"`
	ReadIopsSec      int `json:"read_iops_sec"`
	ReadIopsSecMax   int `json:"read_iops_sec_max"`
	SizeIopsSec      int `json:"size_iops_sec"`
	TotalBytesSec    int `json:"total_bytes_sec"`
	TotalBytesSecMax int `json:"total_bytes_sec_max"`
	TotalIopsSec     int `json:"total_iops_sec"`
	TotalIopsSecMax  int `json:"total_iops_sec_max"`
	WriteBytesSec    int `json:"write_bytes_sec"`
	WriteBytesSecMax int `json:"write_bytes_sec_max"`
	WriteIopsSec     int `json:"write_iops_sec"`
	WriteIopsSecMax  int `json:"write_iops_sec_max"`
}

package compute

// Access Control List
type RecordACL struct {
	// Account ACL list
	AccountACL ListACL `json:"accountAcl"`

	// Compute ACL list
	ComputeACL ListACL `json:"computeAcl"`

	// Resource group ACL list
	RGACL ListACL `json:"rgAcl"`
}

// ACL information
type ItemACL struct {
	// Explicit
	Explicit bool `json:"explicit"`

	// GUID
	GUID string `json:"guid"`

	// Right
	Right string `json:"right"`

	// Status
	Status string `json:"status"`

	// Type
	Type string `json:"type"`

	// User group ID
	UserGroupID string `json:"userGroupId"`
}

// List ACL
type ListACL []ItemACL

// Main information about usage snapshot
type ItemUsageSnapshot struct {
	// Count
	Count uint64 `json:"count,omitempty"`

	// Stored
	Stored float64 `json:"stored"`

	// Label
	Label string `json:"label,omitempty"`

	// Timestamp
	Timestamp uint64 `json:"timestamp,omitempty"`
}

// List of usage snapshot
type ListUsageSnapshots []ItemUsageSnapshot

// Main information about snapshot
type ItemSnapshot struct {
	// List disk ID
	Disks []uint64 `json:"disks"`

	// GUID
	GUID string `json:"guid"`

	// Label
	Label string `json:"label"`

	// Timestamp
	Timestamp uint64 `json:"timestamp"`
}

// List of snapshots
type ListSnapShots []ItemSnapshot

// Main information about port forward
type ItemPFW struct {
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

	// Virtuel machine ID
	VMID uint64 `json:"vmId"`
}

// List port forwards
type ListPFWs []ItemPFW

// Main information about affinity relations
type RecordAffinityRelations struct {
	// Other node
	OtherNode []interface{} `json:"otherNode"`

	// Other node indirect
	OtherNodeIndirect []interface{} `json:"otherNodeIndirect"`

	// Other node indirect soft
	OtherNodeIndirectSoft []interface{} `json:"otherNodeIndirectSoft"`

	// Other node soft
	OtherNodeSoft []interface{} `json:"otherNodeSoft"`

	// Same node
	SameNode []interface{} `json:"sameNode"`

	// Same node soft
	SameNodeSoft []interface{} `json:"sameNodeSoft"`
}

// Main information about attached network
type RecordNetAttach struct {
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

	// Network ID
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

	// List VNF IDs
	VNFs []uint64 `json:"vnfs"`
}

// Detailed information about audit
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

// List Detailed audits
type ListAudits []ItemAudit

// Short information about audit
type ItemShortAudit struct {
	// Epoch
	Epoch float64 `json:"epoch"`

	// Message
	Message string `json:"message"`
}

// List short audits
type ListShortAudits []ItemShortAudit

// Main information about rule
type ItemRule struct {
	// GUID
	GUID string `json:"guid"`

	// Key
	Key string `json:"key"`

	// Mode
	Mode string `json:"mode"`

	// Policy
	Policy string `json:"policy"`

	// Topology
	Topology string `json:"topology"`

	// Value
	Value string `json:"value"`
}

// List rules
type ListRules []ItemRule

// Detailed information about compute
type RecordCompute struct {
	// Access Control List
	ACL RecordACL `json:"ACL"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Affinity label
	AffinityLabel string `json:"affinityLabel"`

	// List affinity rules
	AffinityRules ListRules `json:"affinityRules"`

	// Affinity weight
	AffinityWeight uint64 `json:"affinityWeight"`

	// List anti affinity rules
	AntiAffinityRules ListRules `json:"antiAffinityRules"`

	// Architecture
	Architecture string `json:"arch"`

	// Boot order
	BootOrder []string `json:"bootOrder"`

	// Boot disk size
	BootDiskSize uint64 `json:"bootdiskSize"`

	// Clone reference
	CloneReference uint64 `json:"cloneReference"`

	// List clone IDs
	Clones []uint64 `json:"clones"`

	// Compute CI ID
	ComputeCIID uint64 `json:"computeciId"`

	//  Number of cores
	CPU uint64 `json:"cpus"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Custom fields items
	CustomFields map[string]interface{} `json:"customFields"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// Devices
	Devices interface{} `json:"devices"`

	// List disks in compute
	Disks ListComputeDisks `json:"disks"`

	// Driver
	Driver string `json:"driver"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Image ID
	ImageID uint64 `json:"imageId"`

	// Image name
	ImageName string `json:"imageName"`

	// List interfaces
	Interfaces ListInterfaces `json:"interfaces"`

	// Lock status
	LockStatus string `json:"lockStatus"`

	// Manager ID
	ManagerID uint64 `json:"managerId"`

	// Manager type
	ManagerType string `json:"managerType"`

	// Migration job
	MigrationJob uint64 `json:"migrationjob"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Natable VINS ID
	NatableVINSID uint64 `json:"natableVinsId"`

	// Natable VINS IP
	NatableVINSIP string `json:"natableVinsIp"`

	// Natable VINS Name
	NatableVINSName string `json:"natableVinsName"`

	// Natable VINS network
	NatableVINSNetwork string `json:"natableVinsNetwork"`

	// Natable VINS network name
	NatableVINSNetworkName string `json:"natableVinsNetworkName"`

	// List OS Users
	OSUsers ListOSUser `json:"osUsers"`

	// Pinned or not
	Pinned bool `json:"pinned"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// Reference ID
	ReferenceID string `json:"referenceId"`

	// Registered or not
	Registered bool `json:"registered"`

	// Resource name
	ResName string `json:"resName"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// List snapsets
	SnapSets ListSnapSets `json:"snapSets"`

	// Stateless SepID
	StatelessSepID uint64 `json:"statelessSepId"`

	// Stateless SepType
	StatelessSepType string `json:"statelessSepType"`

	// Status
	Status string `json:"status"`

	// Tags
	Tags map[string]string `json:"tags"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// User Managed or not
	UserManaged bool `json:"userManaged"`

	// Userdata
	Userdata interface{} `json:"userdata"`

	// vGPU IDs
	VGPUs []uint64 `json:"vgpus"`

	// Virtual image ID
	VirtualImageID uint64 `json:"virtualImageId"`

	// Virtual image name
	VirtualImageName string `json:"virtualImageName"`
}

// Main information about OS user
type ItemOSUser struct {
	// GUID
	GUID string `json:"guid"`

	// Login
	Login string `json:"login"`

	// Password
	Password string `json:"password"`

	// Public key
	PubKey string `json:"pubkey"`
}

// List OS users
type ListOSUser []ItemOSUser

// Main information about snapsets
type ItemSnapSet struct {
	// List disk IDs
	Disks []uint64 `json:"disks"`

	// GUID
	GUID string `json:"guid"`

	// Label
	Label string `json:"label"`

	// Timestamp
	Timestamp uint64 `json:"timestamp"`
}

// List snapsets
type ListSnapSets []ItemSnapSet

// Main information about VNF
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

	// Listen SSH or not
	ListenSSH bool `json:"listenSsh"`

	// MAC
	MAC string `json:"mac"`

	// Name
	Name string `json:"name"`

	// Network ID
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

	// List VNF IDs
	VNFs []uint64 `json:"vnfs"`
}

type QOS struct {
	ERate   uint64 `json:"eRate"`
	GUID    string `json:"guid"`
	InBurst uint64 `json:"inBurst"`
	InRate  uint64 `json:"inRate"`
}

// List VNF interfaces
type ListInterfaces []ItemVNFInterface

// List compute disks
type ListComputeDisks []ItemComputeDisk

// Main information about compute disk
type ItemComputeDisk struct {
	// CKey
	CKey string `json:"_ckey"`

	// Access Control List
	ACL map[string]interface{} `json:"acl"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Boot partition
	BootPartition uint64 `json:"bootPartition"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// Destruction time
	DestructionTime uint64 `json:"destructionTime"`

	// Disk path
	DiskPath string `json:"diskPath"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Image ID
	ImageID uint64 `json:"imageId"`

	// List image IDs
	Images []uint64 `json:"images"`

	// IO tune
	IOTune IOTune `json:"iotune"`

	// IQN
	IQN string `json:"iqn"`

	// Login
	Login string `json:"login"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Order
	Order uint64 `json:"order"`

	// Params
	Params string `json:"params"`

	// Parent ID
	ParentID uint64 `json:"parentId"`

	// Password
	Passwd string `json:"passwd"`

	// PCI slot
	PCISlot uint64 `json:"pciSlot"`

	// Pool
	Pool string `json:"pool"`

	// Present to
	PresentTo []uint64 `json:"presentTo"`

	// Purge time
	PurgeTime uint64 `json:"purgeTime"`

	// Reality device number
	RealityDeviceNumber uint64 `json:"realityDeviceNumber"`

	// Resource ID
	ResID string `json:"resId"`

	// Role
	Role string `json:"role"`

	// SepID
	SepID uint64 `json:"sepId"`

	// Shareable
	Shareable bool `json:"shareable"`

	// Size max
	SizeMax uint64 `json:"sizeMax"`

	//Size used
	SizeUsed uint64 `json:"sizeUsed"`

	// List extend snapshots
	Snapshots SnapshotExtendList `json:"snapshots"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`

	// Virtual machine ID
	VMID uint64 `json:"vmid"`
}

// Main information about snapshot extend
type SnapshotExtend struct {
	// GUID
	GUID string `json:"guid"`

	// Label
	Label string `json:"label"`

	// Resource ID
	ResID string `json:"resId"`

	// SnapSetGUID
	SnapSetGUID string `json:"snapSetGuid"`

	// SnapSetTime
	SnapSetTime uint64 `json:"snapSetTime"`

	// TimeStamp
	TimeStamp uint64 `json:"timestamp"`
}

// List Snapshot Extend
type SnapshotExtendList []SnapshotExtend

// Main information about IO tune
type IOTune struct {
	// ReadBytesSec
	ReadBytesSec uint64 `json:"read_bytes_sec"`

	// ReadBytesSecMax
	ReadBytesSecMax uint64 `json:"read_bytes_sec_max"`

	// ReadIOPSSec
	ReadIOPSSec uint64 `json:"read_iops_sec"`

	// ReadIOPSSecMax
	ReadIOPSSecMax uint64 `json:"read_iops_sec_max"`

	// SizeIOPSSec
	SizeIOPSSec uint64 `json:"size_iops_sec"`

	// TotalBytesSec
	TotalBytesSec uint64 `json:"total_bytes_sec"`

	// TotalBytesSecMax
	TotalBytesSecMax uint64 `json:"total_bytes_sec_max"`

	// TotalIOPSSec
	TotalIOPSSec uint64 `json:"total_iops_sec"`

	// TotalIOPSSecMax
	TotalIOPSSecMax uint64 `json:"total_iops_sec_max"`

	// WriteBytesSec
	WriteBytesSec uint64 `json:"write_bytes_sec"`

	// WriteBytesSecMax
	WriteBytesSecMax uint64 `json:"write_bytes_sec_max"`

	// WriteIOPSSec
	WriteIOPSSec uint64 `json:"write_iops_sec"`

	// WriteIOPSSecMax
	WriteIOPSSecMax uint64 `json:"write_iops_sec_max"`
}

// Main information about compute
type ItemCompute struct {
	// Access Control List
	ACL []interface{} `json:"acl"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Affinity label
	AffinityLabel string `json:"affinityLabel"`

	// List affinity rules
	AffinityRules ListRules `json:"affinityRules"`

	// Affinity weight
	AffinityWeight uint64 `json:"affinityWeight"`

	// List anti affinity rules
	AntiAffinityRules ListRules `json:"antiAffinityRules"`

	// Architecture
	Architecture string `json:"arch"`

	// Boot order
	BootOrder []string `json:"bootOrder"`

	// Boot disk size
	BootDiskSize uint64 `json:"bootdiskSize"`

	// Clone reference
	CloneReference uint64 `json:"cloneReference"`

	// List clone IDs
	Clones []uint64 `json:"clones"`

	// Compute CI ID
	ComputeCIID uint64 `json:"computeciId"`

	// Number of cores
	CPU uint64 `json:"cpus"`

	// Created by
	CreatedBy string `json:"createdBy"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Custom fields list
	CustomFields map[string]interface{} `json:"customFields"`

	// Deleted by
	DeletedBy string `json:"deletedBy"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Description
	Description string `json:"desc"`

	// Devices
	Devices interface{} `json:"devices"`

	// List disk items
	Disks []InfoDisk `json:"disks"`

	// Driver
	Driver string `json:"driver"`

	// Grid ID
	GID uint64 `json:"gid"`

	// GUID
	GUID uint64 `json:"guid"`

	// ID
	ID uint64 `json:"id"`

	// Image ID
	ImageID uint64 `json:"imageId"`

	// List interfaces
	Interfaces ListInterfaces `json:"interfaces"`

	// Lock status
	LockStatus string `json:"lockStatus"`

	// Manager ID
	ManagerID uint64 `json:"managerId"`

	// Manager type
	ManagerType string `json:"managerType"`

	// Migration job
	MigrationJob uint64 `json:"migrationjob"`

	// Milestones
	Milestones uint64 `json:"milestones"`

	// Name
	Name string `json:"name"`

	// Pinned or not
	Pinned bool `json:"pinned"`

	// Number of RAM
	RAM uint64 `json:"ram"`

	// Reference ID
	ReferenceID string `json:"referenceId"`

	// Registered
	Registered bool `json:"registered"`

	// Resource name
	ResName string `json:"resName"`

	// Resource group ID
	RGID uint64 `json:"rgId"`

	// Resource group name
	RGName string `json:"rgName"`

	// List snapsets
	SnapSets ListSnapSets `json:"snapSets"`

	// Stateless SepID
	StatelessSepID uint64 `json:"statelessSepId"`

	// Stateless SepType
	StatelessSepType string `json:"statelessSepType"`

	// Status
	Status string `json:"status"`

	// Tags
	Tags map[string]string `json:"tags"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Total disk size
	TotalDiskSize uint64 `json:"totalDisksSize"`

	// Updated by
	UpdatedBy string `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`

	// User Managed or not
	UserManaged bool `json:"userManaged"`

	// List vGPU IDs
	VGPUs []uint64 `json:"vgpus"`

	// VINS connected
	VINSConnected uint64 `json:"vinsConnected"`

	// Virtual image ID
	VirtualImageID uint64 `json:"virtualImageId"`
}

// Information Disk
type InfoDisk struct {
	// ID
	ID uint64 `json:"id"`

	// PCISlot
	PCISlot uint64 `json:"pciSlot"`
}

// List information about computes
type ListComputes []ItemCompute

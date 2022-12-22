package disks

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

// Main information about disks
type InfoDisk struct {
	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Access Control Control
	ACL map[string]interface{} `json:"acl"`

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

	// List of image IDs
	Images []uint64 `json:"images"`

	// IOTune
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
	Password string `json:"passwd"`

	// PCI slot
	PCISlot int64 `json:"pciSlot"`

	// Pool
	Pool string `json:"pool"`

	// Purge attempts
	PurgeAttempts uint64 `json:"purgeAttempts"`

	// Purge time
	PurgeTime uint64 `json:"purgeTime"`

	// Reality device number
	RealityDeviceNumber uint64 `json:"realityDeviceNumber"`

	// Reference ID
	ReferenceID string `json:"referenceId"`

	// Resource ID
	ResID string `json:"resId"`

	// Resource name
	ResName string `json:"resName"`

	// Role
	Role string `json:"role"`

	// SEP ID
	SEPID uint64 `json:"sepId"`

	// Size max
	SizeMax uint64 `json:"sizeMax"`

	// Size used
	SizeUsed uint64 `json:"sizeUsed"`

	// List snapshots
	Snapshots ListSnapshots `json:"snapshots"`

	// Status
	Status string `json:"status"`

	// Tech status
	TechStatus string `json:"techStatus"`

	// Type
	Type string `json:"type"`

	// Virtual machine ID
	VMID uint64 `json:"vmid"`
}

// Detailed indormation about disk
type RecordDisk struct {
	// Device name
	DeviceName string `json:"devicename"`

	// SEP type
	SEPType string `json:"sepType"`

	// Main information about disk
	InfoDisk
}

// Main information for list disks
type ItemDisk struct {
	// Compute ID
	ComputeID uint64 `json:"computeId"`

	// Compute name
	ComputeName string `json:"computeName"`

	// Machine ID
	MachineID uint64 `json:"machineId"`

	// Machine name
	MachineName string `json:"machineName"`

	// Detailed information about disk
	RecordDisk
}

// List disks
type ListDisks []ItemDisk

// Main information about deleted disk
type ItemDeletedDisk struct {
	// Compute ID
	ComputeID uint64 `json:"computeId"`

	// Compute name
	ComputeName string `json:"computeName"`

	// Machine ID
	MachineID uint64 `json:"machineId"`

	// Machine name
	MachineName string `json:"machineName"`

	// Detailed information about disk
	RecordDisk

	// Updated by
	UpdatedBy uint64 `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// List deleted disks
type ListDeletedDisks []ItemDeletedDisk

// Main information about unattached disk
type ItemUnattachedDisk struct {
	// CKey
	CKey string `json:"_ckey"`

	// Meta
	Meta []interface{} `json:"_meta"`

	// Main information about disk
	InfoDisk

	// Updated by
	UpdatedBy uint64 `json:"updatedBy"`

	// Updated time
	UpdatedTime uint64 `json:"updatedTime"`
}

// List unattached disks
type ListUnattachedDisks []ItemUnattachedDisk

// Main information about snapshot
type ItemSnapshot struct {
	// GUID
	GUID string `json:"guid"`

	// Label
	Label string `json:"label"`

	// Resource ID
	ResID string `json:"resId"`

	// SnapSet GUID
	SnapSetGUID string `json:"snapSetGuid"`

	// SnapSet time
	SnapSetTime uint64 `json:"snapSetTime"`

	// Timestamp
	Timestamp uint64 `json:"timestamp"`
}

// List snapshots
type ListSnapshots []ItemSnapshot

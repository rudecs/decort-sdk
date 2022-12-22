package disks

// Main information about disk
type ItemDisk struct {
	// Access Control List
	ACL map[string]interface{} `json:"acl"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Boot partition
	BootPartition uint64 `json:"bootPartition"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Compute ID
	ComputeID uint64 `json:"computeId"`

	// Compute name
	ComputeName string `json:"computeName"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Device name
	DeviceName string `json:"devicename"`

	// Description
	Description string `json:"desc"`

	// Destruction time
	DestructionTime uint64 `json:"destructionTime"`

	// Grid ID
	GID uint64 `json:"gid"`

	// ID
	ID uint64 `json:"id"`

	// Image ID
	ImageID uint64 `json:"imageId"`

	// List of image IDs
	Images []uint64 `json:"images"`

	// IOTune
	IOTune IOTune `json:"iotune"`

	// Machine ID
	MachineID uint64 `json:"machineId"`

	// Machine name
	MachineName string `json:"machineName"`

	// Name
	Name string `json:"name"`

	// Order
	Order uint64 `json:"order"`

	// Params
	Params string `json:"params"`

	// Parent ID
	ParentID uint64 `json:"parentId"`

	// PCI slot
	PCISlot uint64 `json:"pciSlot"`

	// Pool
	Pool string `json:"pool"`

	// Purge time
	PurgeTime uint64 `json:"purgeTime"`

	// Resource ID
	ResID string `json:"resId"`

	// Resource name
	ResName string `json:"resName"`

	// Role
	Role string `json:"role"`

	// SepType
	SepType string `json:"sepType"`

	// SepID
	SepID uint64 `json:"sepId"`

	// Size max
	SizeMax uint64 `json:"sizeMax"`

	// List of snapshots
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

// List of disks
type ListDisks []ItemDisk

// Main information about snapshot
type ItemSnapshot struct {
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

// List of snapshots
type ListSnapshots []ItemSnapshot

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

// Detailed information about disk
type RecordDisk struct {
	// Access Control List
	ACL map[string]interface{} `json:"acl"`

	// Account ID
	AccountID uint64 `json:"accountId"`

	// Account name
	AccountName string `json:"accountName"`

	// Created time
	CreatedTime uint64 `json:"createdTime"`

	// Deleted time
	DeletedTime uint64 `json:"deletedTime"`

	// Device name
	DeviceName string `json:"devicename"`

	// Description
	Description string `json:"desc"`

	// Destruction time
	DestructionTime uint64 `json:"destructionTime"`

	// Grid ID
	GID uint64 `json:"gid"`

	// ID
	ID uint64 `json:"id"`

	// Image ID
	ImageID uint64 `json:"imageId"`

	// List of image IDs
	Images []uint64 `json:"images"`

	// IOTune
	IOTune IOTune `json:"iotune"`

	// Name
	Name string `json:"name"`

	// Order
	Order uint64 `json:"order"`

	// Params
	Params string `json:"params"`

	// Parent ID
	ParentID uint64 `json:"parentId"`

	// PCI slot
	PCISlot uint64 `json:"pciSlot"`

	// Pool
	Pool string `json:"pool"`

	// Purge time
	PurgeTime uint64 `json:"purgeTime"`

	// Resource ID
	ResID string `json:"resId"`

	// Resource name
	ResName string `json:"resName"`

	// Role
	Role string `json:"role"`

	// SepType
	SepType string `json:"sepType"`

	// SepID
	SepID uint64 `json:"sepId"`

	// Size max
	SizeMax uint64 `json:"sizeMax"`

	// Size used
	SizeUsed uint64 `json:"sizeUsed"`

	// List of snapshots
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

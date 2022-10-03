package disks

type IOTune struct {
	ReadBytesSec     uint64 `json:"read_bytes_sec"`
	ReadBytesSecMax  uint64 `json:"read_bytes_sec_max"`
	ReadIOPSSec      uint64 `json:"read_iops_sec"`
	ReadIOPSSecMax   uint64 `json:"read_iops_sec_max"`
	SizeIOPSSec      uint64 `json:"size_iops_sec"`
	TotalBytesSec    uint64 `json:"total_bytes_sec"`
	TotalBytesSecMax uint64 `json:"total_bytes_sec_max"`
	TotalIOPSSec     uint64 `json:"total_iops_sec"`
	TotalIOPSSecMax  uint64 `json:"total_iops_sec_max"`
	WriteBytesSec    uint64 `json:"write_bytes_sec"`
	WriteBytesSecMax uint64 `json:"write_bytes_sec_max"`
	WriteIOPSSec     uint64 `json:"write_iops_sec"`
	WriteIOPSSecMax  uint64 `json:"write_iops_sec_max"`
}

type Info struct {
	AccountID           uint64                 `json:"accountId"`
	AccountName         string                 `json:"accountName"`
	ACL                 map[string]interface{} `json:"acl"`
	BootPartition       uint64                 `json:"bootPartition"`
	CreatedTime         uint64                 `json:"createdTime"`
	DeletedTime         uint64                 `json:"deletedTime"`
	Desc                string                 `json:"desc"`
	DestructionTime     uint64                 `json:"destructionTime"`
	DiskPath            string                 `json:"diskPath"`
	GID                 uint64                 `json:"gid"`
	GUID                uint64                 `json:"guid"`
	ID                  uint64                 `json:"id"`
	ImageID             uint64                 `json:"imageId"`
	Images              []uint64               `json:"images"`
	IOTune              IOTune                 `json:"iotune"`
	Iqn                 string                 `json:"iqn"`
	Login               string                 `json:"login"`
	Milestones          uint64                 `json:"milestones"`
	Name                string                 `json:"name"`
	Order               uint64                 `json:"order"`
	Params              string                 `json:"params"`
	ParentID            uint64                 `json:"parentId"`
	Passwd              string                 `json:"passwd"`
	PCISlot             int64                  `json:"pciSlot"`
	Pool                string                 `json:"pool"`
	PurgeAttempts       uint64                 `json:"purgeAttempts"`
	PurgeTime           uint64                 `json:"purgeTime"`
	RealityDeviceNumber uint64                 `json:"realityDeviceNumber"`
	ReferenceID         string                 `json:"referenceId"`
	ResID               string                 `json:"resId"`
	ResName             string                 `json:"resName"`
	Role                string                 `json:"role"`
	SepID               uint64                 `json:"sepId"`
	SizeMax             uint64                 `json:"sizeMax"`
	SizeUsed            uint64                 `json:"sizeUsed"`
	Snapshots           []Snapshot             `json:"snapshots"`
	Status              string                 `json:"status"`
	TechStatus          string                 `json:"techStatus"`
	Type                string                 `json:"type"`
	VMID                uint64                 `json:"vmid"`
}

type Disk struct {
	DeviceName string `json:"devicename"`
	SepType    string `json:"sepType"`
	Info
}

type ListDisks []struct {
	ComputeID   uint64 `json:"computeId"`
	ComputeName string `json:"computeName"`
	MachineID   uint64 `json:"machineId"`
	MachineName string `json:"machineName"`
	Disk
}

type ListDeletedDisks []struct {
	ComputeID   uint64 `json:"computeId"`
	ComputeName string `json:"computeName"`
	MachineID   uint64 `json:"machineId"`
	MachineName string `json:"machineName"`
	Disk
	UpdatedBy   uint64 `json:"updatedBy"`
	UpdatedTime uint64 `json:"updatedTime"`
}

type ListUnattached []struct {
	CKey string        `json:"_ckey"`
	Meta []interface{} `json:"_meta"`
	Info
	UpdatedBy   uint64 `json:"updatedBy"`
	UpdatedTime uint64 `json:"updatedTime"`
}

type Snapshot struct {
	GUID        string `json:"guid"`
	Label       string `json:"label"`
	ResID       string `json:"resId"`
	SnapSetGUID string `json:"snapSetGuid"`
	SnapSetTime uint64 `json:"snapSetTime"`
	Timestamp   uint64 `json:"timestamp"`
}

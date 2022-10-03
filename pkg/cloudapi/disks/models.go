package disks

type Disk struct {
	ACL             map[string]interface{} `json:"acl"`
	AccountID       uint64                 `json:"accountId"`
	AccountName     string                 `json:"accountName"`
	BootPartition   uint64                 `json:"bootPartition"`
	CreatedTime     uint64                 `json:"createdTime"`
	ComputeID       uint64                 `json:"computeId"`
	ComputeName     string                 `json:"computeName"`
	DeletedTime     uint64                 `json:"deletedTime"`
	DeviceName      string                 `json:"devicename"`
	Description     string                 `json:"desc"`
	DestructionTime uint64                 `json:"destructionTime"`
	GID             uint64                 `json:"gid"`
	ID              uint64                 `json:"id"`
	ImageID         uint64                 `json:"imageId"`
	Images          []uint64               `json:"images"`
	IOTune          IOTune                 `json:"iotune"`
	MachineID       uint64                 `json:"machineId"`
	MachineName     string                 `json:"machineName"`
	Name            string                 `json:"name"`
	Order           uint64                 `json:"order"`
	Params          string                 `json:"params"`
	ParentID        uint64                 `json:"parentId"`
	PciSlot         uint64                 `json:"pciSlot"`
	Pool            string                 `json:"pool"`
	PurgeTime       uint64                 `json:"purgeTime"`
	ResID           string                 `json:"resId"`
	ResName         string                 `json:"resName"`
	Role            string                 `json:"role"`
	SepType         string                 `json:"sepType"`
	SepID           uint64                 `json:"sepId"` // NOTE: absent from compute/get output
	SizeMax         uint64                 `json:"sizeMax"`
	Snapshots       []Snapshot             `json:"snapshots"`
	Status          string                 `json:"status"`
	TechStatus      string                 `json:"techStatus"`
	Type            string                 `json:"type"`
	VMID            uint64                 `json:"vmid"`
}

type Snapshot struct {
	Guid        string `json:"guid"`
	Label       string `json:"label"`
	ResID       string `json:"resId"`
	SnapSetGuid string `json:"snapSetGuid"`
	SnapSetTime uint64 `json:"snapSetTime"`
	TimeStamp   uint64 `json:"timestamp"`
}

type SnapshotList []Snapshot

type IOTune struct {
	ReadBytesSec     uint64 `json:"read_bytes_sec"`
	ReadBytesSecMax  uint64 `json:"read_bytes_sec_max"`
	ReadIopsSec      uint64 `json:"read_iops_sec"`
	ReadIopsSecMax   uint64 `json:"read_iops_sec_max"`
	SizeIopsSec      uint64 `json:"size_iops_sec"`
	TotalBytesSec    uint64 `json:"total_bytes_sec"`
	TotalBytesSecMax uint64 `json:"total_bytes_sec_max"`
	TotalIopsSec     uint64 `json:"total_iops_sec"`
	TotalIopsSecMax  uint64 `json:"total_iops_sec_max"`
	WriteBytesSec    uint64 `json:"write_bytes_sec"`
	WriteBytesSecMax uint64 `json:"write_bytes_sec_max"`
	WriteIopsSec     uint64 `json:"write_iops_sec"`
	WriteIopsSecMax  uint64 `json:"write_iops_sec_max"`
}

type DiskList []Disk

type DisksTypesListCoomon []string

type DisksTypesListDetailed struct {
	Pools []Pool `json:"pools"`
	SepID uint64 `json:"sepId"`
}

type Pool struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

type DiskRecord struct {
	ACL             map[string]interface{} `json:"acl"`
	AccountID       uint64                 `json:"accountId"`
	AccountName     string                 `json:"accountName"`
	CreatedTime     uint64                 `json:"createdTime"`
	DeletedTime     uint64                 `json:"deletedTime"`
	DeviceName      string                 `json:"devicename"`
	Description     string                 `json:"desc"`
	DestructionTime uint64                 `json:"destructionTime"`
	GID             uint64                 `json:"gid"`
	ID              uint64                 `json:"id"`
	ImageID         uint64                 `json:"imageId"`
	Images          []uint64               `json:"images"`
	IOTune          IOTune                 `json:"iotune"`
	Name            string                 `json:"name"`
	Order           uint64                 `json:"order"`
	Params          string                 `json:"params"`
	ParentID        uint64                 `json:"parentId"`
	PciSlot         uint64                 `json:"pciSlot"`
	Pool            string                 `json:"pool"`
	PurgeTime       uint64                 `json:"purgeTime"`
	ResID           string                 `json:"resId"`
	ResName         string                 `json:"resName"`
	Role            string                 `json:"role"`
	SepType         string                 `json:"sepType"`
	SepID           uint64                 `json:"sepId"` // NOTE: absent from compute/get output
	SizeMax         uint64                 `json:"sizeMax"`
	SizeUsed        uint64                 `json:"sizeUsed"` // sum over all snapshots of this disk to report total consumed space
	Snapshots       []Snapshot             `json:"snapshots"`
	Status          string                 `json:"status"`
	TechStatus      string                 `json:"techStatus"`
	Type            string                 `json:"type"`
	VMID            uint64                 `json:"vmid"`
}

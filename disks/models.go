package disks

type Disk struct {
	Acl             map[string]interface{} `json:"acl"`
	AccountID       int                    `json:"accountId"`
	AccountName     string                 `json:"accountName"`
	BootPartition   int                    `json:"bootPartition"`
	CreatedTime     uint64                 `json:"createdTime"`
	ComputeID       int                    `json:"computeId"`
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
	MachineId       int                    `json:"machineId"`
	MachineName     string                 `json:"machineName"`
	Name            string                 `json:"name"`
	Order           int                    `json:"order"`
	Params          string                 `json:"params"`
	ParentId        uint64                 `json:"parentId"`
	PciSlot         uint64                 `json:"pciSlot"`
	Pool            string                 `json:"pool"`
	PurgeTime       uint64                 `json:"purgeTime"`
	ResID           string                 `json:"resId"`
	ResName         string                 `json:"resName"`
	Role            string                 `json:"role"`
	SepType         string                 `json:"sepType"`
	SepID           int                    `json:"sepId"` // NOTE: absent from compute/get output
	SizeMax         int                    `json:"sizeMax"`
	Snapshots       []Snapshot             `json:"snapshots"`
	Status          string                 `json:"status"`
	TechStatus      string                 `json:"techStatus"`
	Type            string                 `json:"type"`
	VMID            int                    `json:"vmid"`
}

type Snapshot struct {
	Guid        string `json:"guid"`
	Label       string `json:"label"`
	ResId       string `json:"resId"`
	SnapSetGuid string `json:"snapSetGuid"`
	SnapSetTime uint64 `json:"snapSetTime"`
	TimeStamp   uint64 `json:"timestamp"`
}

type SnapshotList []Snapshot

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

type DiskList []Disk

type DisksTypesListCoomon []string

type DisksTypesListDetailed struct {
	Pools []Pool `json:"pools"`
	SepId uint64 `json:"sepId"`
}

type Pool struct {
	Name  string   `json:"name"`
	Types []string `json:"types"`
}

type DiskRecord struct {
	Acl             map[string]interface{} `json:"acl"`
	AccountID       int                    `json:"accountId"`
	AccountName     string                 `json:"accountName"`
	CreatedTime     uint64                 `json:"createdTime"`
	DeletedTime     uint64                 `json:"deletedTime"`
	DeviceName      string                 `json:"devicename"`
	Description     string                 `json:"desc"`
	DestructionTime uint64                 `json:"destructionTime"`
	GID             int                    `json:"gid"`
	ID              uint                   `json:"id"`
	ImageID         int                    `json:"imageId"`
	Images          []int                  `json:"images"`
	IOTune          IOTune                 `json:"iotune"`
	Name            string                 `json:"name"`
	Order           int                    `json:"order"`
	Params          string                 `json:"params"`
	ParentId        int                    `json:"parentId"`
	PciSlot         int                    `json:"pciSlot"`
	Pool            string                 `json:"pool"`
	PurgeTime       uint64                 `json:"purgeTime"`
	ResID           string                 `json:"resId"`
	ResName         string                 `json:"resName"`
	Role            string                 `json:"role"`
	SepType         string                 `json:"sepType"`
	SepID           int                    `json:"sepId"` // NOTE: absent from compute/get output
	SizeMax         int                    `json:"sizeMax"`
	SizeUsed        int                    `json:"sizeUsed"` // sum over all snapshots of this disk to report total consumed space
	Snapshots       []Snapshot             `json:"snapshots"`
	Status          string                 `json:"status"`
	TechStatus      string                 `json:"techStatus"`
	Type            string                 `json:"type"`
	VMID            int                    `json:"vmid"`
}

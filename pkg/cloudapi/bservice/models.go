package bservice

type BasicService struct {
	AccountID     uint64     `json:"accountId"`
	AccountName   string     `json:"accountName"`
	BaseDomain    string     `json:"baseDomain"`
	Computes      []Compute  `json:"computes"`
	CPUTotal      uint64     `json:"cpuTotal"`
	CreatedBy     string     `json:"createdBy"`
	CreatedTime   uint64     `json:"createdTime"`
	DeletedBy     string     `json:"deletedBy"`
	DeletedTime   uint64     `json:"deletedTime"`
	DiskTotal     uint64     `json:"diskTotal"`
	GID           uint64     `json:"gid"`
	Groups        []uint64   `json:"groups"`
	GroupsName    []string   `json:"groupsName"`
	GUID          uint64     `json:"guid"`
	ID            uint64     `json:"id"`
	Milestones    uint64     `json:"milestones"`
	Name          string     `json:"name"`
	ParentSrvID   uint64     `json:"parentSrvId"`
	ParentSrvType string     `json:"parentSrvType"`
	RAMTotal      uint64     `json:"ramTotal"`
	RGID          uint64     `json:"rgId"`
	RGName        string     `json:"rgName"`
	Snapshots     []Snapshot `json:"snapshots"`
	SSHKey        string     `json:"sshKey"`
	SSHUser       string     `json:"sshUser"`
	Status        string     `json:"status"`
	TechStatus    string     `json:"techStatus"`
	UpdatedBy     string     `json:"updatedBy"`
	UpdatedTime   uint64     `json:"updatedTime"`
	UserManaged   bool       `json:"userManaged"`
}

type Compute struct {
	CompGroupID   uint64 `json:"compgroupId"`
	CompGroupName string `json:"compgroupName"`
	CompGroupRole string `json:"compgroupRole"`
	ID            uint64 `json:"id"`
	Name          string `json:"name"`
}

type Snapshot struct {
	GUID      string `json:"guid"`
	Label     string `json:"label"`
	Timestamp uint64 `json:"timestamp"`
	Valid     bool   `json:"valid"`
}

type Group struct {
	AccountID    uint64         `json:"accountId"`
	AccountName  string         `json:"accountName"`
	Computes     []GroupCompute `json:"computes"`
	Consistency  bool           `json:"consistency"`
	CPU          uint64         `json:"cpu"`
	CreatedBy    string         `json:"createdBy"`
	CreatedTime  uint64         `json:"createdTime"`
	DeletedBy    string         `json:"deletedBy"`
	DeletedTime  uint64         `json:"deletedTime"`
	Disk         uint64         `json:"disk"`
	Driver       string         `json:"driver"`
	Extnets      []uint64       `json:"extnets"`
	GID          uint64         `json:"gid"`
	GUID         uint64         `json:"guid"`
	ID           uint64         `json:"id"`
	ImageID      uint64         `json:"imageId"`
	Milestones   uint64         `json:"milestones"`
	Name         string         `json:"name"`
	Parents      []uint64       `json:"parents"`
	RAM          uint64         `json:"ram"`
	RGID         uint64         `json:"rgId"`
	RGName       string         `json:"rgName"`
	Role         string         `json:"role"`
	SepID        uint64         `json:"sepId"`
	SeqNo        uint64         `json:"seqNo"`
	ServiceID    uint64         `json:"serviceId"`
	Status       string         `json:"status"`
	TechStatus   string         `json:"techStatus"`
	TimeoutStart uint64         `json:"timeoutStart"`
	UpdatedBy    string         `json:"updatedBy"`
	UpdatedTime  uint64         `json:"updatedTime"`
	VINSes       []uint64       `json:"vinses"`
}

type GroupCompute struct {
	ID          uint64   `json:"id"`
	IPAddresses []string `json:"ipAddresses"`
	Name        string   `json:"name"`
	OSUsers     []OSUser `json:"osUsers"`
}

type OSUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type BasicServiceShort struct {
	AccountID     uint64   `json:"accountId"`
	AccountName   string   `json:"accountName"`
	BaseDomain    string   `json:"baseDomain"`
	CreatedBy     string   `json:"createdBy"`
	CreatedTime   uint64   `json:"createdTime"`
	DeletedBy     string   `json:"deletedBy"`
	DeletedTime   uint64   `json:"deletedTime"`
	GID           uint64   `json:"gid"`
	Groups        []uint64 `json:"groups"`
	GUID          uint64   `json:"guid"`
	ID            uint64   `json:"id"`
	Name          string   `json:"name"`
	ParentSrvID   uint64   `json:"parentSrvId"`
	ParentSrvType string   `json:"parentSrvType"`
	RGID          uint64   `json:"rgId"`
	RGName        string   `json:"rgName"`
	SSHUser       string   `json:"sshUser"`
	Status        string   `json:"status"`
	TechStatus    string   `json:"techStatus"`
	UpdatedBy     string   `json:"updatedBy"`
	UpdatedTime   uint64   `json:"updatedTime"`
	UserManaged   bool     `json:"userManaged"`
}

type BasicServiceList []BasicServiceShort

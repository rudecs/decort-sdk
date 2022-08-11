package k8s

type K8SGroup struct {
	Annotations  []string         `json:"annotations"`
	CPU          uint             `json:"cpu"`
	DetailedInfo DetailedInfoList `json:"detailedInfo"`
	Disk         uint             `json:"disk"`
	GUID         string           `json:"guid"`
	ID           uint64           `json:"id"`
	Labels       []string         `json:"labels"`
	Name         string           `json:"name"`
	Num          uint             `json:"num"`
	RAM          uint             `json:"ram"`
	Taints       []string         `json:"taints"`
}

type K8SGroupList []K8SGroup

type DetailedInfo struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	TechStatus string `json:"techStatus"`
}

type DetailedInfoList []DetailedInfo

type K8SRecord struct {
	ACL         ACLGroup  `json:"ACL"`
	AccountId   uint64    `json:"accountId"`
	AccountName string    `json:"accountName"`
	BServiceId  uint64    `json:"bserviceId"`
	CIId        uint64    `json:"ciId"`
	CreatedBy   string    `json:"createdBy"`
	CreatedTime uint64    `json:"createdTime"`
	DeletedBy   string    `json:"deletedBy"`
	DeletedTime uint64    `json:"deletedTime"`
	ID          uint64    `json:"id"`
	K8CIName    string    `json:"k8ciName"`
	K8SGroups   K8SGroups `json:"k8sGroups"`
	LBId        uint64    `json:"lbId"`
	Name        string    `json:"name"`
	RGID        uint64    `json:"rgId"`
	RGName      string    `json:"rgName"`
	Status      string    `json:"status"`
	TechStatus  string    `json:"techStatus"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedTime uint64    `json:"updatedTime"`
}

type K8SGroups struct {
	Masters MasterGroup  `json:"masters"`
	Workers K8SGroupList `json:"workers"`
}

type MasterGroup struct {
	CPU          uint             `json:"cpu"`
	DetailedInfo DetailedInfoList `json:"detailedInfo"`
	Disk         uint             `json:"disk"`
	ID           uint64           `json:"id"`
	Name         string           `json:"name"`
	Num          uint             `json:"num"`
	RAM          uint             `json:"ram"`
}

type ACLGroup struct {
	AccountAcl AclList `json:"accountAcl"`
	K8SAcl     AclList `json:"k8sAcl"`
	RGAcl      AclList `json:"rgAcl"`
}

type Acl struct {
	Explicit    bool   `json:"explicit"`
	GUID        string `json:"guid"`
	Right       string `json:"right"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	UserGroupId string `json:"userGroupId"`
}

type AclList []Acl

type K8SItem struct {
	AccountId      uint64         `json:"accountId"`
	AccountName    string         `json:"accountName"`
	Acl            string         `json:"acl"`
	BServiceId     uint64         `json:"bserviceId"`
	CIId           uint64         `json:"ciId"`
	Config         string         `json:"config"`
	CreatedBy      string         `json:"createdBy"`
	CreatedTime    uint64         `json:"createdTime"`
	DeletedBy      string         `json:"deletedBy"`
	DeletedTime    uint64         `json:"deletedTime"`
	Description    string         `json:"desc"`
	ExtnetId       uint64         `json:"extnetId"`
	GID            uint64         `json:"gid"`
	GUID           uint64         `json:"guid"`
	ID             uint64         `json:"id"`
	LBId           uint64         `json:"lbId"`
	Milestones     uint64         `json:"milestones"`
	Name           string         `json:"name"`
	RGID           uint64         `json:"rgId"`
	RGName         string         `json:"rgName"`
	ServiceAccount ServiceAccount `json:"serviceAccount"`
	Status         string         `json:"status"`
	TechStatus     string         `json:"techStatus"`
	UpdatedBy      string         `json:"updatedBy"`
	UpdatedTime    uint64         `json:"updatedTime"`
	VinsId         uint64         `json:"vinsId"`
	WorkersGroup   K8SGroupList   `json:"workersGroups"`
}

type ServiceAccount struct {
	GUID     string `json:"guid"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type K8SList []K8SItem

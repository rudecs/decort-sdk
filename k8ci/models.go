package k8ci

type K8CIItem struct {
	CreatedTime uint64 `json:"createdTime"`
	K8CIRecord
}

type K8CIList []K8CIItem

type K8CIRecord struct {
	Description string `json:"desc"`
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
}

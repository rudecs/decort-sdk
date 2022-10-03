package sizes

type Size struct {
	Description string   `json:"desc"`
	Disks       []uint64 `json:"disks"`
	ID          uint64   `json:"id"`
	Memory      uint64   `json:"memory"`
	Name        string   `json:"name"`
	VCPUs       uint64   `json:"vcpus"`
}

type SizesList []Size

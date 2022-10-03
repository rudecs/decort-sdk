package disks

import (
	"github.com/rudecs/decort-sdk/interfaces"
)

type Disks struct {
	client interfaces.Caller
}

func New(client interfaces.Caller) *Disks {
	return &Disks{
		client,
	}
}

package client

import (
	"github.com/rudecs/decort-sdk/disks"
)

func (dc *decortClient) Disks() *disks.Disks {
	return disks.New(dc)
}

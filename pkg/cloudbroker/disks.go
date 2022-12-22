package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/disks"
)

// Accessing the Disks method group
func (cb *CloudBroker) Disks() *disks.Disks {
	return disks.New(cb.client)
}

package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/disks"
)

func (cb *CloudBroker) Disks() *disks.Disks {
	return disks.New(cb.client)
}

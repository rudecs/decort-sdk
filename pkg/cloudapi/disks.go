package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/disks"
)

// Accessing the Disks method group
func (ca *CloudAPI) Disks() *disks.Disks {
	return disks.New(ca.client)
}

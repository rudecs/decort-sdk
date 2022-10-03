package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/disks"
)

func (ca *CloudApi) Disks() *disks.Disks {
	return disks.New(ca.client)
}

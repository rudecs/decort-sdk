package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/disks"
)

func (dc *decortClient) Disks() *disks.Disks {
	return disks.New(dc)
}

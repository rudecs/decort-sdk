package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/disks"
)

func (dc *Client) Disks() *disks.Disks {
	return disks.New(dc)
}

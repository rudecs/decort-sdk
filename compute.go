package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/compute"
)

func (dc *Client) Compute() *compute.Compute {
	return compute.New(dc)
}

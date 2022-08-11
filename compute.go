package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/compute"
)

func (dc *decortClient) Compute() *compute.Compute {
	return compute.New(dc)
}

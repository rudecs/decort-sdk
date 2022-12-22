package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/compute"
)

// Accessing the Compute method group
func (ca *CloudAPI) Compute() *compute.Compute {
	return compute.New(ca.client)
}

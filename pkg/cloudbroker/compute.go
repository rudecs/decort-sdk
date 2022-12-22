package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/compute"
)

// Accessing the Compute method group
func (cb *CloudBroker) Compute() *compute.Compute {
	return compute.New(cb.client)
}

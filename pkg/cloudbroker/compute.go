package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/compute"
)

func (cb *CloudBroker) Compute() *compute.Compute {
	return compute.New(cb.client)
}

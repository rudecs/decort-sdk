package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/computeci"
)

func (cb *CloudBroker) ComputeCI() *computeci.ComputeCI {
	return computeci.New(cb.client)
}

package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudbroker/rg"

// Accessing the RG method group
func (cb *CloudBroker) RG() *rg.RG {
	return rg.New(cb.client)
}

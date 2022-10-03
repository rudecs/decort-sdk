package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudapi/rg"

func (cb *CloudBroker) RG() *rg.RG {
	return rg.New(cb.client)
}

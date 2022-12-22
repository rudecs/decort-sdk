package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudbroker/sep"

// Accessing the SEP method group
func (cb *CloudBroker) SEP() *sep.SEP {
	return sep.New(cb.client)
}

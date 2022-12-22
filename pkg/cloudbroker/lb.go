package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudbroker/lb"

// Accessing the LB method group
func (cb *CloudBroker) LB() *lb.LB {
	return lb.New(cb.client)
}

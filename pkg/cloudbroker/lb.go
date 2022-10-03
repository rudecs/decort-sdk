package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudapi/lb"

func (cb *CloudBroker) LB() *lb.LB {
	return lb.New(cb.client)
}

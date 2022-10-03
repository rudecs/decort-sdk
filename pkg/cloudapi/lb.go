package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/lb"

func (ca *CloudApi) LB() *lb.LB {
	return lb.New(ca.client)
}

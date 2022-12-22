package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/lb"

// Accessing the LB method group
func (ca *CloudAPI) LB() *lb.LB {
	return lb.New(ca.client)
}

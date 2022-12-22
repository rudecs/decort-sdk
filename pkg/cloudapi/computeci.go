package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/computeci"
)

// Accessing the ComputeCI method group
func (ca *CloudAPI) ComputeCI() *computeci.ComputeCI {
	return computeci.New(ca.client)
}

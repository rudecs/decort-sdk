package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/computeci"
)

func (ca *CloudApi) ComputeCI() *computeci.ComputeCI {
	return computeci.New(ca.client)
}

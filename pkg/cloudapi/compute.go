package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/compute"
)

func (ca *CloudApi) Compute() *compute.Compute {
	return compute.New(ca.client)
}

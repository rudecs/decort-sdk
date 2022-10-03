package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/flipgroup"
)

func (ca *CloudApi) FlipGroup() *flipgroup.FlipGroup {
	return flipgroup.New(ca.client)
}

package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/flipgroup"
)

func (cb *CloudBroker) FlipGroup() *flipgroup.FlipGroup {
	return flipgroup.New(cb.client)
}

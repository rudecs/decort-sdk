package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/sizes"
)

func (cb *CloudBroker) Sizes() *sizes.Sizes {
	return sizes.New(cb.client)
}

package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/sizes"
)

func (ca *CloudApi) Sizes() *sizes.Sizes {
	return sizes.New(ca.client)
}

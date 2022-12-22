package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/sizes"
)

// Accessing the Sizes method group
func (ca *CloudAPI) Sizes() *sizes.Sizes {
	return sizes.New(ca.client)
}

package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/flipgroup"
)

// Accessing the FLIPGroup method group
func (ca *CloudAPI) FLIPGroup() *flipgroup.FLIPGroup {
	return flipgroup.New(ca.client)
}

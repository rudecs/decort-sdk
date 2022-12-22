package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/vins"
)

// Accessing the VINS method group
func (ca *CloudAPI) VINS() *vins.VINS {
	return vins.New(ca.client)
}

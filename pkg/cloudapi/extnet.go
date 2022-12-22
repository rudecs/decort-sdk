package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/extnet"
)

// Accessing the ExtNet method group
func (ca *CloudAPI) ExtNet() *extnet.ExtNet {
	return extnet.New(ca.client)
}

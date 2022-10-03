package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/extnet"
)

func (ca *CloudApi) ExtNet() *extnet.ExtNet {
	return extnet.New(ca.client)
}

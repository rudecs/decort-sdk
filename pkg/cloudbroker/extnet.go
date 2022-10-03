package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/extnet"
)

func (cb *CloudBroker) ExtNet() *extnet.ExtNet {
	return extnet.New(cb.client)
}

package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/extnet"
)

// Accessing the ExtNet method group
func (cb *CloudBroker) ExtNet() *extnet.ExtNet {
	return extnet.New(cb.client)
}

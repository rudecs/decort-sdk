package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/vins"
)

// Accessing the VINS method group
func (cb *CloudBroker) VINS() *vins.VINS {
	return vins.New(cb.client)
}

package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/vins"
)

func (cb *CloudBroker) VINS() *vins.VINS {
	return vins.New(cb.client)
}

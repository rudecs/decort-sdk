package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/kvmx86"
)

func (cb *CloudBroker) KVMX86() *kvmx86.KVMX86 {
	return kvmx86.New(cb.client)
}

package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/kvmx86"
)

func (ca *CloudApi) KVMX86() *kvmx86.KVMX86 {
	return kvmx86.New(ca.client)
}

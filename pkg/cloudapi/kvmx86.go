package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/kvmx86"
)

// Accessing the KVMX86 method group
func (ca *CloudAPI) KVMX86() *kvmx86.KVMX86 {
	return kvmx86.New(ca.client)
}

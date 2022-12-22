package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8ci"
)

// Accessing the K8CI method group
func (ca *CloudAPI) K8CI() *k8ci.K8CI {
	return k8ci.New(ca.client)
}

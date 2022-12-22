package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/k8ci"
)

// Accessing the K8CI method group
func (cb *CloudBroker) K8CI() *k8ci.K8CI {
	return k8ci.New(cb.client)
}

package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8ci"
)

func (cb *CloudBroker) K8CI() *k8ci.K8CI {
	return k8ci.New(cb.client)
}

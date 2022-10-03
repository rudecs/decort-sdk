package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8ci"
)

func (ca *CloudApi) K8CI() *k8ci.K8CI {
	return k8ci.New(ca.client)
}

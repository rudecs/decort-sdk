package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8ci"
)

func (dc *decortClient) K8CI() *k8ci.K8CI {
	return k8ci.New(dc)
}
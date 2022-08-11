package client

import (
	"github.com/rudecs/decort-sdk/k8ci"
)

func (dc *decortClient) K8CI() *k8ci.K8CI {
	return k8ci.New(dc)
}

package client

import (
	"github.com/rudecs/decort-sdk/k8s"
)

func (dc *decortClient) K8S() *k8s.K8S {
	return k8s.New(dc)
}

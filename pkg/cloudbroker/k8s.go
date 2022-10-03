package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8s"
)

func (cb *CloudBroker) K8S() *k8s.K8S {
	return k8s.New(cb.client)
}

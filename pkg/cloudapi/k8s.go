package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8s"
)

func (ca *CloudApi) K8S() *k8s.K8S {
	return k8s.New(ca.client)
}

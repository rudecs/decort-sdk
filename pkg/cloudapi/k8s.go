package cloudapi

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8s"
)

// Accessing the K8S method group
func (ca *CloudAPI) K8S() *k8s.K8S {
	return k8s.New(ca.client)
}

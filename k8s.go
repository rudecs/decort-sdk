package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/k8s"
)

func (dc *decortClient) K8S() *k8s.K8S {
	return k8s.New(dc)
}

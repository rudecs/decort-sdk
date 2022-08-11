package decortsdk

import "github.com/rudecs/decort-sdk/pkg/cloudapi/kvmppc"

func (dc *decortClient) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(dc)
}

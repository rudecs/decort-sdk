package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudapi/kvmppc"

func (cb *CloudBroker) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(cb.client)
}

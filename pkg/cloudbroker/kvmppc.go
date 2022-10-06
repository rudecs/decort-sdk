package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudbroker/kvmppc"

func (cb *CloudBroker) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(cb.client)
}

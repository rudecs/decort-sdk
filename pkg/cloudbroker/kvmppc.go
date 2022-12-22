package cloudbroker

import "github.com/rudecs/decort-sdk/pkg/cloudbroker/kvmppc"

// Accessing the KVMPPC method group
func (cb *CloudBroker) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(cb.client)
}

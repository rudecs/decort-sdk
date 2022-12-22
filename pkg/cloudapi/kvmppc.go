package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/kvmppc"

// Accessing the KVMPPC method group
func (ca *CloudAPI) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(ca.client)
}

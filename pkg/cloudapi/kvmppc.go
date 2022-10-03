package cloudapi

import "github.com/rudecs/decort-sdk/pkg/cloudapi/kvmppc"

func (ca *CloudApi) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(ca.client)
}

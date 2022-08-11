package client

import "github.com/rudecs/decort-sdk/kvmppc"

func (dc *decortClient) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(dc)
}

package decortsdk

import "github.com/rudecs/decort-sdk/pkg/cloudapi/kvmppc"

func (dc *Client) KVMPPC() *kvmppc.KVMPPC {
	return kvmppc.New(dc)
}

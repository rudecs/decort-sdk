package decortsdk

import "github.com/rudecs/decort-sdk/pkg/cloudapi/lb"

func (dc *Client) LB() *lb.LB {
	return lb.New(dc)
}

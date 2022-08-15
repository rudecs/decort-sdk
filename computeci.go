package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/computeci"
)

func (dc *Client) ComputeCI() *computeci.ComputeCI {
	return computeci.New(dc)
}

package decortsdk

import (
	"github.com/rudecs/decort-sdk/pkg/cloudapi/flipgroup"
)

func (dc *Client) FlipGroup() *flipgroup.FlipGroup {
	return flipgroup.New(dc)
}

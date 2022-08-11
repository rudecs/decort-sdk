package client

import (
	"github.com/rudecs/decort-sdk/flipgroup"
)

func (dc *decortClient) FlipGroup() *flipgroup.FlipGroup {
	return flipgroup.New(dc)
}

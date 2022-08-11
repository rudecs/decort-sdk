package client

import (
	"github.com/rudecs/decort-sdk/sizes"
)

func (dc *decortClient) Sizes() *sizes.Sizes {
	return sizes.New(dc)
}

package client

import (
	"github.com/rudecs/decort-sdk/compute"
)

func (dc *decortClient) Compute() *compute.Compute {
	return compute.New(dc)
}

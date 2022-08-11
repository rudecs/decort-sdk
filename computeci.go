package client

import (
	"github.com/rudecs/decort-sdk/computeci"
)

func (dc *decortClient) ComputeCI() *computeci.ComputeCI {
	return computeci.New(dc)
}

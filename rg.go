package client

import "github.com/rudecs/decort-sdk/rg"

func (dc *decortClient) RG() *rg.RG {
	return rg.New(dc)
}

package client

import (
	"github.com/rudecs/decort-sdk/vins"
)

func (dc *decortClient) Vins() *vins.Vins {
	return vins.New(dc)
}

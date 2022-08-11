package client

import (
	"github.com/rudecs/decort-sdk/extnet"
)

func (dc *decortClient) Extnet() *extnet.Extnet {
	return extnet.New(dc)
}

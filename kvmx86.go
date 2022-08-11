package client

import (
	"github.com/rudecs/decort-sdk/kvmx86"
)

func (dc *decortClient) KVMX86() *kvmx86.KVMX86 {
	return kvmx86.New(dc)
}

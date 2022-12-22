package cloudbroker

import (
	"github.com/rudecs/decort-sdk/pkg/cloudbroker/image"
)

// Accessing the Image method group
func (cb *CloudBroker) Image() *image.Image {
	return image.New(cb.client)
}
